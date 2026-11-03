//ff:func feature=inspect type=decoder control=sequence
//ff:what TestNewEntry — words include the TAG WORD and the name is the core record name (a Value ISSUER format by format name), or a slot marker's name
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestNewEntry(t *testing.T) {
	e := newEntry(core.EncodeUintValue(5))
	if len(e.Words) != 2 || e.Words[1] != 5 || e.Name != "UNSIGNED-INTEGER" || e.Data != nil {
		t.Fatalf("got %+v", e)
	}
	if e := newEntry(core.Record{Kind: core.KindValue, Codec: core.CodecValueIssuer2Word, Issuer: []core.Word{0, 2}, Format: 0x0702, Payload: make([]core.Word, 8)}); e.Name != "RATIONAL-I64" {
		t.Fatalf("issuer format: %+v", e)
	}
	if e := newEntry(core.Record{Kind: core.KindReference, Codec: 2}); e.Name != "NO-VALUE" || len(e.Words) != 1 {
		t.Fatalf("marker: %+v", e)
	}
	if e := newEntry(core.Record{Kind: core.KindReference, Codec: 1, Payload: []core.Word{1}}); e.Name != "" {
		t.Fatalf("Codec 0000 with payload: %+v", e)
	}
}
