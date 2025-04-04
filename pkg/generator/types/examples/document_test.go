package examples

import (
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/schema/extensions"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	builder := NewDocumentBuilder(extensions.ReadmeExample{
		Language: YAML.String(),
	}, 0)

	builder.Add("first", "benoit")
	builder.Add("address", NewObject("address"))
	builder.Add("town", "{St Hil}")
	builder.Add("zip", 38660)
	builder.Pop()
	builder.Add("last", "Bordigoni")
	builder.Add("kids", NewArray("kids"))
	builder.Add("", NewObject(""))
	builder.Add("firstname", "Eliott")
	builder.Pop()
	builder.Add("", NewObject(""))
	builder.Add("firstname", "Nino")
	builder.Pop()
	builder.Pop()
	builder.Add("works", NewArray("works"))
	builder.Pop()

	json, err := builder.Marshall()
	if err != nil {
		return
	}
	fmt.Println(json)
	if !strings.Contains(json, "St Hil") {
		panic("error")
	}

}
