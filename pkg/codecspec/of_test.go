//ff:func feature=codecspec type=helper control=sequence
//ff:what TestOf — a declared codec's fields; undeclared and ISSUER records are false
package codecspec

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestOf(t *testing.T) {
	s := Spec{0x0980: {{Name: "doc"}}}
	if fs, ok := s.Of(core.Record{Kind: core.KindEdge, Codec: core.EdgePublication}); !ok || fs[0].Name != "doc" {
		t.Fatalf("got %+v %v", fs, ok)
	}
	if _, ok := s.Of(core.Record{Kind: core.KindEdge, Codec: core.EdgeDeprecation}); ok {
		t.Fatal("undeclared: want false")
	}
	if _, ok := s.Of(core.Record{Kind: core.KindEdge, Codec: core.EdgePublication, Issuer: []core.Word{1}}); ok {
		t.Fatal("issuer: want false")
	}
}
