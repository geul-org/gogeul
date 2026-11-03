//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestCheckIssuerDocPointer — R5: ISSUER-, DOC- and CODEC-POINTER Issuers of 2/4/16 WORD in canonical width pass; the list ID F…F under Issuer 0x00000000 passes; 1 WORD, odd widths, zero-padded Issuers and F…F under another Issuer fail
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCheckIssuerDocPointer(t *testing.T) {
	ref := func(c uint16, p ...core.Word) rec {
		return recOf(core.Record{Kind: core.KindReference, Codec: c, Payload: p})
	}
	iss, doc, cdc := core.CodecIssuerPointer, core.CodecDocPointer, core.CodecCodecPointer
	cases := []struct {
		r    rec
		c    uint16
		want string
	}{
		{ref(iss, 0, 0x0D), iss, ""},
		{ref(iss, 1, 2), iss, ""},
		{ref(iss, 0x40), iss, "R5"},
		{ref(iss, 0, 0, 1, 2), iss, "R5"},
		{ref(iss, 0, 1, 0, 0), iss, ""},
		{ref(iss, 1, 2, 3), iss, "R5"},
		{ref(iss), iss, "R5"},
		{ref(doc, 0, 0, 0xFFFF, 0, 0, 1, 0, 1), doc, ""},
		{ref(doc, 0, 0, 0xFFFF, 0, 0, 1, 0, 0), doc, ""},
		{ref(doc, 0, 0, 0, 0, 0, 0, 0, 1), doc, ""},
		{ref(doc, 0, 0, 0xFFFF, 0xFFFF, 0xFFFF, 0xFFFF, 0, 1), doc, ""},
		{ref(doc, 0, 2, 0xFFFF, 0xFFFF, 0xFFFF, 0xFFFF, 0, 1), doc, "R5"},
		{ref(doc, 0, 0xFFFF, 0, 0, 1, 0, 1), doc, "R5"},
		{ref(doc, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1), doc, "R5"},
		{ref(cdc, 0, 0, 0, 0x0B80, 0, 0), cdc, ""},
		{ref(cdc, 0, 1, 0, 0, 0, 0x0B80, 0, 1), cdc, ""},
		{ref(cdc, 0, 0, 0, 1, 0, 0x0B80, 0, 1), cdc, "R5"},
		{ref(cdc, 0, 0, 0x0B80, 0, 1), cdc, "R5"},
	}
	for i, tc := range cases {
		if got := rulesOf(checkIssuerDocPointer(tc.r, tc.c)); got != tc.want {
			t.Fatalf("%d: %q want %q", i, got, tc.want)
		}
	}
}
