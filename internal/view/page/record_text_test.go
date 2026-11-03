//ff:func feature=page type=helper control=iteration dimension=1
//ff:what TestRecordText — the Absent record as —, the other slot markers by name, references by number, decoded Values (text, language text, URI, EDTF, numbers, FRACTION), undecodable ones as record name (ISSUER format name) and hex, a lone TAG by name
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestRecordText(t *testing.T) {
	for want, r := range map[string]core.Record{
		"#3":                  core.EncodeLocalRef(3),
		"—":                   {Kind: core.KindReference},
		`"Douglas Adams"@en`:  core.EncodeLanguageText(core.EncodeBytesValue(core.CodecLanguageBCP47Even, []byte("en")), []byte("Douglas Adams")),
		"codec 0x1e 0x4142":   {Kind: core.KindValue, Codec: 0x1E, Payload: []core.Word{0x4142}},
		"https://x":           core.EncodeBytesValue(core.CodecURIEven, []byte("https://x")),
		"1974":                core.EncodeBytesValue(core.CodecEDTFEven, []byte("1974")),
		"7":                   core.EncodeUintValue(7),
		"UTF8-TEXT 0x0001":    {Kind: core.KindValue, Codec: core.CodecUTF8TextEven | 1, Payload: []core.Word{1}},
		"PUBLICATION":         {Kind: core.KindEdge, Codec: 0x18},
		"UNKNOWN-VALUE":       {Kind: core.KindReference, Codec: 1},
		"NO-VALUE":            {Kind: core.KindReference, Codec: 2},
		"2/3":                 {Kind: core.KindValue, Codec: core.CodecFraction, Payload: []core.Word{0x0411, 2, 0x0421, 3}},
		"RATIONAL-I64 0x0001": {Kind: core.KindValue, Codec: core.CodecValueIssuer2Word, Issuer: []core.Word{0, 2}, Format: 0x0702, Payload: []core.Word{1}},
	} {
		if got := recordText(r); got != want {
			t.Fatalf("got %q want %q", got, want)
		}
	}
}
