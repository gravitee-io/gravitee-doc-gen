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

package util

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Template", func() {
	When("using level 2 markdown", func() {

		It("generates a level 3 content", func() {

			given := `
## Header 1
Lorem ipsum dolor sit amet
### Subsection 1.1
Lorem ipsum dolor sit amet
### Subsection 1.2
Lorem ipsum dolor sit amet
## Header 2
Lorem ipsum dolor sit amet
### Subsection 2.1
Lorem ipsum dolor sit amet
`

			expectedLevel3 := `
### Header 1
Lorem ipsum dolor sit amet
#### Subsection 1.1
Lorem ipsum dolor sit amet
#### Subsection 1.2
Lorem ipsum dolor sit amet
### Header 2
Lorem ipsum dolor sit amet
#### Subsection 2.1
Lorem ipsum dolor sit amet
`
			result := util.MoveMarkdownHeader(1, given)
			Expect(result).To(Equal(expectedLevel3))
		})

	})
})

func TestMdHeaderShift(t *testing.T) {

}

var _ = Describe("Template functions", func() {
	render := func(content string) (string, error) {
		file := filepath.Join(GinkgoT().TempDir(), "template.tmpl")
		Expect(os.WriteFile(file, []byte(content), 0o600)).To(Succeed())
		rendered, err := util.RenderTemplateFromFile(file, nil)
		return string(rendered), err
	}

	It("tells whether a file exists", func() {
		file := filepath.Join(GinkgoT().TempDir(), "present.yaml")
		Expect(os.WriteFile(file, []byte("a: b"), 0o600)).To(Succeed())

		rendered, err := render(`{{ fileExists "` + file + `" }} {{ fileExists "missing.yaml" }}`)
		Expect(err).NotTo(HaveOccurred())
		Expect(rendered).To(Equal("true false"))
	})

	It("fails the rendering with the given message", func() {
		_, err := render(`before {{ fail "move the examples" }}`)
		Expect(err).To(MatchError(ContainSubstring("move the examples")))
	})
})
