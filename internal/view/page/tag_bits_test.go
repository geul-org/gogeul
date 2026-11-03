//ff:func feature=page type=helper control=iteration dimension=1
//ff:what TestTagBits — prefix, KKK and Kind, C and its name (5Q Quantifier, slot marker, ISSUER), LLLL and its length; fallback and hex for unnamed codes
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestTagBits(t *testing.T) {
	cases := map[core.Word]string{
		0x1BD9: "000 Proposal\n110 EDGE\n111101 EDGE-ISSUER-2WORD\n1001 10 WORDs",
		0x0051: "000 Proposal\n000 REFERENCE\n000101 LOCAL · Identified\n0001 2 WORDs",
		0x0181: "000 Proposal\n000 REFERENCE\n011000 WIKIDATA-QID · Variable\n0001 2 WORDs",
		0x1C10: "000 Proposal\n111 META\n000001 DOC-2WORD\n0000 1 WORD",
		0x0E1E: "000 Proposal\n011 EDGE\n100001 0x061\n1110 Length 확장 1 WORD",
		0x0E6F: "000 Proposal\n011 EDGE\n100110 0x066\n1111 Length 확장 2 WORD",
		0x0010: "000 Proposal\n000 REFERENCE\n000001 자리 표지 UNKNOWN-VALUE\n0000 1 WORD",
		0x03C6: "000 Proposal\n000 REFERENCE\n111100 CODEC-POINTER\n0110 7 WORDs",
	}
	for w, want := range cases {
		if got := tagBits(w, ""); got != want {
			t.Fatalf("%04X: %q want %q", uint16(w), got, want)
		}
	}
	if got := tagBits(0x0E38, "RULE-DECLARE"); got != "000 Proposal\n011 EDGE\n100011 RULE-DECLARE\n1000 9 WORDs" {
		t.Fatalf("fallback: %q", got)
	}
}
