// Copyright (C) 2015 The Gravitee team (http://gravitee.io)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package output

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"maps"
	"os"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/bootstrap"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/chunks"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/config"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
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
	data := util.Unstructured{}
	for _, chunk := range generated {
		data[chunk.ID] = chunk
	}

	// make bootstrap data available
	maps.Copy(data, bootstrap.GetExported())

	// render template
	if rendered, err := util.RenderTemplateFromFile(output.Template, data); err == nil {
		// add to buffer
		buffer.Write(rendered)
	} else {
		return fmt.Errorf("error rendering main template: %w", err)
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
