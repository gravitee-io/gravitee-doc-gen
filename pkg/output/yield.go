package output

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/chunks"
	"github.com/gravitee-io-labs/readme-gen/pkg/config"
	"github.com/gravitee-io-labs/readme-gen/pkg/util"
	"io"
	"log"
	"os"
)

const generatedFromMarker = "<!-- generated-start -->"
const readmeFileName = "README.md"

func Yield(cfg config.Config, pl config.Plugin, generated []chunks.Generated, write bool) error {

	buf := make([]byte, 0)
	buffer := bytes.NewBuffer(buf)

	// Read README.md Keep what is before marker
	buffer.Write(readUntilMarker())
	// add marker for next generation
	buffer.WriteString("\n")
	buffer.WriteString(generatedFromMarker)
	buffer.WriteString("\n")

	// chunks to map
	data := make(map[string]any)
	for _, chunk := range generated {
		data[chunk.Id] = chunk
	}
	data[config.PluginChunkId] = pl

	// render template
	if rendered, err := util.RenderTemplateFromFile(cfg.MainTemplate, data); err == nil {
		// add to buffer
		buffer.Write(rendered)
	} else {
		return errors.New(fmt.Sprintf("Error rendering main template: %s", err))
	}

	// write buffer
	_, err := buffer.WriteTo(chooseWriter(write))

	return err
}

func chooseWriter(write bool) io.Writer {
	if write {
		return File{}
	}
	return Console{}
}

func readUntilMarker() []byte {
	file, err := os.Open(readmeFileName)
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
		if bytes.Contains(b, []byte(generatedFromMarker)) {
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
