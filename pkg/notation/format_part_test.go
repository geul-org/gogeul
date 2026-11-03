//ff:func feature=notation type=decoder control=sequence
//ff:what TestFormatPart — a record part, an ENUM name, raw words
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestFormatPart(t *testing.T) {
	if formatPart(inspect.Part{Words: []core.Word{0x0051, 2}, Record: true}) != "#2" || formatPart(inspect.Part{Words: []core.Word{1}, Enum: "normal"}) != "normal" || formatPart(inspect.Part{Words: []core.Word{1}}) != "0x0001" {
		t.Fatal("parts")
	}
}
