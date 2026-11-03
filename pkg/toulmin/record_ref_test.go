//ff:func feature=toulmin type=decoder control=iteration dimension=1
//ff:what TestRecordRef — direct and GEUL-format Edge records resolve; other kinds and Issuers fail
package toulmin

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestRecordRef(t *testing.T) {
	if r, err := RecordRef(core.Record{Kind: core.KindEdge, Codec: CodecRun}); err != nil || r.Codec != CodecRun || r.Format != 0 {
		t.Fatalf("direct: %+v %v", r, err)
	}
	if r, err := RecordRef(core.Record{Kind: core.KindEdge, Codec: core.CodecEdgeIssuer2Word, Issuer: []core.Word{0, 0}, Format: FormatText}); err != nil || r.Format != FormatText {
		t.Fatalf("format: %+v %v", r, err)
	}
	bad := []core.Record{
		{Kind: core.KindValue},
		{Kind: core.KindEdge, Codec: core.CodecEdgeIssuer2Word, Issuer: []core.Word{0, 1}},
		{Kind: core.KindEdge, Codec: core.CodecEdgeIssuer2Word, Issuer: []core.Word{1, 0}},
		{Kind: core.KindEdge, Codec: core.CodecEdgeIssuer4Word, Issuer: []core.Word{0, 1, 0, 0}},
	}
	for i, b := range bad {
		if _, err := RecordRef(b); err == nil {
			t.Fatalf("bad %d: want error", i)
		}
	}
}
