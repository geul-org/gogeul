//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestCheckLanguageText — valid language and text (a leading U+FEFF too) pass; a missing or wrong language record, a bad language record, bad UTF-8 and bad padding fail
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCheckLanguageText(t *testing.T) {
	en := core.EncodeBytesValue(core.CodecLanguageBCP47Even, []byte("en"))
	v := func(r core.Record) rec { return recOf(r) }
	cases := []struct {
		r    core.Record
		want string
	}{
		{core.EncodeLanguageText(en, []byte("DNA")), ""},
		{core.EncodeLanguageText(en, []byte("\ufeffAbecásia (GE)")), ""},
		{core.EncodeLanguageText(core.Record{Kind: core.KindValue, Codec: core.CodecLanguageShortCode, Payload: []core.Word{7}}, []byte("가")), ""},
		{core.Record{Kind: core.KindValue, Codec: core.CodecUTF8LanguageTextEven, Payload: []core.Word{0x056F}}, "V7"},
		{core.EncodeLanguageText(core.EncodeUintValue(5), []byte("x")), "V7"},
		{core.EncodeLanguageText(core.EncodeBytesValue(core.CodecLanguageBCP47Even, nil), []byte("x")), "V5"},
		{core.EncodeLanguageText(en, []byte("\xff\xfe")), "V5"},
		{core.Record{Kind: core.KindValue, Codec: core.CodecUTF8LanguageTextEven | 1, Payload: []core.Word{0x0561, 0x656E, 0x4142}}, "V4"},
	}
	for i, c := range cases {
		if got := rulesOf(checkLanguageText(v(c.r))); got != c.want {
			t.Fatalf("case %d: %q want %q", i, got, c.want)
		}
	}
}
