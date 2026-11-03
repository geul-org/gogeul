//ff:func feature=page type=helper control=iteration dimension=1
//ff:what TestRefPayloadLines — EXTERNAL-ISSUER-1WORD payload names Issuer, 문서 ID and 대상 ID at their first words
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestRefPayloadLines(t *testing.T) {
	r := core.Record{Kind: core.KindReference, Codec: core.CodecExternalIssuer2Word, Payload: []core.Word{0, 0, 0xffff, 0, 0, 2, 0}}
	lines := refPayloadLines(r)
	names := []string{"Issuer", "", "문서 ID", "", "", "", "대상 ID"}
	for i, l := range lines {
		if l.Name != names[i] {
			t.Fatalf("line %d: %+v", i, l)
		}
	}
	if len(lines) != 7 || lines[2].Word != "FFFF" {
		t.Fatalf("got %+v", lines)
	}
}
