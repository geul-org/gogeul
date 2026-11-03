//ff:func feature=page type=helper control=iteration dimension=1
//ff:what TestTagBits — prefix, section KKKK and Kind, C and its name (5Q Quantifier in either Reference section, slot marker, ISSUER), LLLL and its length; fallback and hex for unnamed codes
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestTagBits(t *testing.T) {
	cases := map[core.Word]string{
		0x1BD9: "000 Proposal\n1101 EDGE\n11101 EDGE-ISSUER-2WORD\n1001 10 WORDs",
		0x0051: "000 Proposal\n0000 REFERENCE\n00101 LOCAL · Identified\n0001 2 WORDs",
		0x0181: "000 Proposal\n0000 REFERENCE\n11000 WIKIDATA-QID · Variable\n0001 2 WORDs",
		0x0241: "000 Proposal\n0001 REFERENCE\n00100 EXTERNAL-POINTER-1WORD · Variable\n0001 2 WORDs",
		0x1C10: "000 Proposal\n1110 META\n00001 DOC-2WORD\n0000 1 WORD",
		0x101E: "000 Proposal\n1000 EDGE\n00001 0x061\n1110 Length 확장 1 WORD",
		0x106F: "000 Proposal\n1000 EDGE\n00110 0x066\n1111 Length 확장 2 WORD",
		0x0010: "000 Proposal\n0000 REFERENCE\n00001 자리 표지 UNKNOWN-VALUE\n0000 1 WORD",
		0x0426: "000 Proposal\n0010 REFERENCE\n00010 CODEC-POINTER\n0110 7 WORDs",
		0x0790: "000 Proposal\n0011 VALUE\n11001 UTF8-LANGUAGE-TEXT\n0000 1 WORD",
	}
	for w, want := range cases {
		if got := tagBits(w, ""); got != want {
			t.Fatalf("%04X: %q want %q", uint16(w), got, want)
		}
	}
	if got := tagBits(0x1038, "RULE-DECLARE"); got != "000 Proposal\n1000 EDGE\n00011 RULE-DECLARE\n1000 9 WORDs" {
		t.Fatalf("fallback: %q", got)
	}
}
