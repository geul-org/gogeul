//ff:func feature=toulmin type=encoder control=sequence
//ff:what TestEdgeRecord — direct codecs and GEUL format numbers become Edge records
package toulmin

import (
	"reflect"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestEdgeRecord(t *testing.T) {
	if got, err := EdgeRecord(CodecRef{Codec: CodecAttacks}, []core.Word{1}); err != nil || !reflect.DeepEqual(got, []core.Word{0x0E81, 1}) {
		t.Fatalf("direct: %04X %v", got, err)
	}
	if got, err := EdgeRecord(CodecRef{Codec: core.CodecEdgeIssuer2Word, Format: FormatText}, []core.Word{1}); err != nil || !reflect.DeepEqual(got, []core.Word{0x1BD5, 0, 0, 0, 0xC001, 1}) {
		t.Fatalf("format: %04X %v", got, err)
	}
}
