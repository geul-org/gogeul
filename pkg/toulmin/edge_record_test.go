//ff:func feature=toulmin type=encoder control=sequence
//ff:what TestEdgeRecord — direct codecs and GEUL format numbers become Edge records
package toulmin

import (
	"reflect"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestEdgeRecord(t *testing.T) {
	if got, err := EdgeRecord(CodecRef{Codec: CodecAttacks}, []core.Word{1}); err != nil || !reflect.DeepEqual(got, []core.Word{0x1081, 1}) {
		t.Fatalf("direct: %04X %v", got, err)
	}
	if got, err := EdgeRecord(CodecRef{Codec: CodecText}, []core.Word{1}); err != nil || !reflect.DeepEqual(got, []core.Word{0x1211, 1}) {
		t.Fatalf("moved code: %04X %v", got, err)
	}
}
