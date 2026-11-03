//ff:func feature=toulmin type=decoder control=iteration dimension=1
//ff:what TestRecordRef — direct Edge records (TOULMIN and the §3.14 moved codes) resolve; other kinds and every EDGE-ISSUER record, the old GEUL formats included, fail
package toulmin

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestRecordRef(t *testing.T) {
	if r, err := RecordRef(core.Record{Kind: core.KindEdge, Codec: CodecRun}); err != nil || r.Codec != CodecRun {
		t.Fatalf("direct: %+v %v", r, err)
	}
	if r, err := RecordRef(core.Record{Kind: core.KindEdge, Codec: CodecText}); err != nil || r.Codec != CodecText {
		t.Fatalf("moved direct code: %+v %v", r, err)
	}
	bad := []core.Record{
		{Kind: core.KindValue},
		{Kind: core.KindEdge, Codec: core.CodecEdgeIssuer2Word, Issuer: []core.Word{0, 0}, Format: 0xC001},
		{Kind: core.KindEdge, Codec: core.CodecEdgeIssuer2Word, Issuer: []core.Word{0, 1}},
		{Kind: core.KindEdge, Codec: core.CodecEdgeIssuer2Word, Issuer: []core.Word{1, 2}},
		{Kind: core.KindEdge, Codec: core.CodecEdgeIssuer4Word, Issuer: []core.Word{0, 1, 0, 0}},
	}
	for i, b := range bad {
		if _, err := RecordRef(b); err == nil {
			t.Fatalf("bad %d: want error", i)
		}
	}
}
