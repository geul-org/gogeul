//ff:func feature=notation type=decoder control=iteration dimension=1
//ff:what TestFormatValue — each short form (integers, FRACTION included); broken payloads and codecs without one are false
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestFormatValue(t *testing.T) {
	lang := core.EncodeBytesValue(core.CodecLanguageBCP47Even, []byte("en"))
	sd, _ := core.EncodeDecimalValue(true, -2, 120)
	q, _ := core.EncodeQuantity(core.Quantity{Amount: sd, Unit: absent(), Lower: absent(), Upper: absent()})
	cases := map[string]core.Record{
		`"a"`:                                    core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("a")),
		`"Adams"@en`:                             core.EncodeLanguageText(lang, []byte("Adams")),
		"1.20":                                   sd,
		"UNSIGNED-INTEGER(5)":                    core.EncodeUintValue(5),
		"POSIX-TIMESTAMP-SECONDS(9)":             core.EncodePosixSeconds(9),
		"QUANTITY(1.20, ABSENT, ABSENT, ABSENT)": q,
		"SHA1-DIGEST(0x0001)":                    {Kind: core.KindValue, Codec: core.CodecSHA1Digest, Payload: []core.Word{1}},
		"SIGNED-INTEGER(-1)":                     {Kind: core.KindValue, Codec: core.CodecSignedInteger, Payload: []core.Word{0xFFFF}},
		"FRACTION(2, 3)":                         {Kind: core.KindValue, Codec: core.CodecFraction, Payload: []core.Word{0x0611, 2, 0x0621, 3}},
	}
	for want, r := range cases {
		if got, ok := formatValue(r); !ok || got != want {
			t.Errorf("%s: got %s %v", want, got, ok)
		}
	}
	bad := []core.Record{
		{Kind: core.KindValue, Codec: core.CodecUTF8TextEven | 1},
		{Kind: core.KindValue, Codec: core.CodecUnsignedInteger, Payload: []core.Word{0, 1}},
		{Kind: core.KindValue, Codec: core.CodecBinaryFloat, Payload: []core.Word{0}},
	}
	for _, r := range bad {
		if s, ok := formatValue(r); ok {
			t.Errorf("%+v: want false, got %s", r, s)
		}
	}
}
