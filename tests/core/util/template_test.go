package util

import (
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
