//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestCheckText — UTF-8 for text (leading or inner U+FEFF passes), ASCII for URI, language tags, EDTF and GEDCOM-DATE (spaces inside), plain-day EDTF fails V6
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCheckText(t *testing.T) {
	cases := []struct {
		even uint16
		b    string
		want string
	}{
		{core.CodecUTF8TextEven, "가", ""},
		{core.CodecUTF8TextEven, "\xff", "V5"},
		{core.CodecUTF8TextEven, "\xef\xbb\xbfx", ""},
		{core.CodecUTF8TextEven, "\xef\xbb\xbf", ""},
		{core.CodecUTF8TextEven, "x\xef\xbb\xbf", ""},
		{core.CodecUTF8TextEven, "\xef\xbb\xbf\xff", "V5"},
		{core.CodecURIEven, "http://x", ""},
		{core.CodecURIEven, "", "V5"},
		{core.CodecLanguageBCP47Even, "ko KR", "V5"},
		{core.CodecEDTFEven, "1974", ""},
		{core.CodecEDTFEven, "197X", ""},
		{core.CodecEDTFEven, "1952-03-11?", ""},
		{core.CodecEDTFEven, "", "V5"},
		{core.CodecEDTFEven, "1952-03-11", "V6"},
		{core.CodecEDTFEven, "-0043-03-15", "V6"},
		{core.CodecGEDCOMDateEven, "JULIAN 14 OCT 1066", ""},
		{core.CodecGEDCOMDateEven, "JULIAN 1300 ", "V5"},
		{core.CodecGEDCOMDateEven, "", "V5"},
	}
	for i, c := range cases {
		if got := rulesOf(checkText(rec{}, c.even, []byte(c.b))); got != c.want {
			t.Fatalf("case %d: %q", i, got)
		}
	}
}
