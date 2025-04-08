package output

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/bootstrap"
	"github.com/gravitee-io-labs/readme-gen/pkg/chunks"
	"github.com/gravitee-io-labs/readme-gen/pkg/config"
	"github.com/gravitee-io-labs/readme-gen/pkg/util"
	"io"
	"log"
	"os"
)

const generatedMarker = "<!-- GENERATED CODE - DO NOT ALTER THIS OR THE FOLLOWING LINES -->"

func Yield(output config.Output, generated []chunks.Generated, write bool) error {

	buffer := bytes.Buffer{}

	if output.ProcessExisting {
		// Read target file with what is before the marker
		buffer.Write(readUntilMarker(output.Target))
		// add marker for next generation
		buffer.WriteString("\n")
		buffer.WriteString(generatedMarker)
		buffer.WriteString("\n")
	}

	// chunks to map
	data := make(map[string]any)
	for _, chunk := range generated {
		data[chunk.Id] = chunk
	}

	for key, export := range bootstrap.Registry.GetExports() {
		data[export] = bootstrap.Registry.GetData(key)
	}

	// render template
	if rendered, err := util.RenderTemplateFromFile(output.Template, data); err == nil {
		// add to buffer
		buffer.Write(rendered)
	} else {
		return errors.New(fmt.Sprintf("Error rendering main template: %s", err))
	}

	// write buffer
	_, err := buffer.WriteTo(chooseWriter(write, output.Target))

	return err
}

func chooseWriter(write bool, file string) io.Writer {
	if write {
		return File{To: file}
	}
	return Console{To: file}
}

func readUntilMarker(target string) []byte {
	file, err := os.Open(target)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = file.Close()
	}()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	header := make([]byte, 0)
	found := false
	for scanner.Scan() {
		b := scanner.Bytes()
		if bytes.Contains(b, []byte(generatedMarker)) {
			found = true
			break
		} else {
			header = append(header, b...)
		}
	}
	if !found {
		return make([]byte, 0)
	}

	return header
}
