//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestS314Direct — the 25 former GEUL ISSUER formats are direct codes: right Kind, not ISSUER, codec ID = TAG with L=0, Value names (pair stripped), Edge left to pkg/toulmin, a TAG then Payload record
package core

import (
	"slices"
	"strings"
	"testing"
)

func TestS314Direct(t *testing.T) {
	rows := loadS314(t)
	if len(rows) != 25 {
		t.Fatalf("%d rows, want 25", len(rows))
	}
	for _, r := range rows {
		k, c := SplitTarget(r.Tag)
		if k.String() != strings.ToUpper(r.Kind) || IssuerWidth(k, c) != 0 || uint32(TargetWord(k, c)) != r.CodecID {
			t.Fatalf("%s: %v %X", r.Name, k, c)
		}
		name := strings.TrimSuffix(strings.TrimSuffix(r.Name, "-EVEN"), "-ODD")
		if k == KindValue && CodecName(k, c) != name || k == KindEdge && (CodecName(k, c) != "" || !coreUnnamedEdge(c)) {
			t.Fatalf("%s: name %q", r.Name, CodecName(k, c))
		}
		ws, err := EncodeRecord(Record{Kind: k, Codec: c, Payload: []Word{7}})
		if err != nil || !slices.Equal(ws, []Word{r.Tag | 1, 7}) {
			t.Fatalf("%s: %04X %v", r.Name, ws, err)
		}
		if got := IssuerFormatName(Record{Kind: KindValue, Codec: CodecValueIssuer2Word, Issuer: []Word{0, 0}, Format: r.Format}); got != "" {
			t.Fatalf("%s: GEUL ISSUER format still named %q", r.Name, got)
		}
	}
}
