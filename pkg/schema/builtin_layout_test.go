//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestBuiltinLayout — each core-embedded Edge with a bad payload fails; good ones and other codecs pass
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestBuiltinLayout(t *testing.T) {
	for _, c := range []uint16{core.EdgePublication, core.EdgeDeclare, core.EdgeField, core.EdgeEnum, core.EdgeCodec} {
		if want, err := builtinLayout(core.Record{Kind: core.KindEdge, Codec: c}); err == nil || want == "" {
			t.Fatalf("0x%02X: want error", c)
		}
	}
	good := core.EncodeFieldDecl(core.FieldDecl{EdgeID: 1, Width: 2})
	if _, err := builtinLayout(good); err != nil {
		t.Fatalf("good: %v", err)
	}
	if _, err := builtinLayout(core.Record{Kind: core.KindEdge, Codec: 0x05}); err != nil {
		t.Fatalf("other: %v", err)
	}
}
