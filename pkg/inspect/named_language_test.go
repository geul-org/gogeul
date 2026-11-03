//ff:func feature=inspect type=decoder control=sequence
//ff:what TestNamedLanguage — a short code the registry has becomes its tag; an unknown code, a BCP 47 tag, another Value and a malformed record stay as ParseValue gives them
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestNamedLanguage(t *testing.T) {
	langs := map[string]string{"0x0784": "en", "0x4130": "pt-BR"}
	short := core.EncodeLanguageText(core.Record{Kind: core.KindValue, Codec: core.CodecLanguageShortCode, Payload: []core.Word{0x4130}}, []byte("Brasil"))
	if v, ok := namedLanguage(short, nil); !ok || v.Lang != "0x4130" {
		t.Fatalf("no registry: %+v %v", v, ok)
	}
	if v, ok := namedLanguage(short, langs); !ok || v.Lang != "pt-BR" || v.Text != "Brasil" {
		t.Fatalf("named: %+v %v", v, ok)
	}
	unknown := core.EncodeLanguageText(core.Record{Kind: core.KindValue, Codec: core.CodecLanguageShortCode, Payload: []core.Word{0x7000}}, []byte("x"))
	if v, _ := namedLanguage(unknown, langs); v.Lang != "0x7000" {
		t.Fatalf("unknown code: %q", v.Lang)
	}
	bcp := core.EncodeLanguageText(core.EncodeBytesValue(core.CodecLanguageBCP47Even, []byte("0x0784")), []byte("x"))
	if v, _ := namedLanguage(bcp, langs); v.Lang != "0x0784" {
		t.Fatalf("a BCP 47 language is never renamed: %q", v.Lang)
	}
	if v, ok := namedLanguage(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("0x0784")), langs); !ok || v.Text != "0x0784" || v.Lang != "" {
		t.Fatalf("other Value: %+v", v)
	}
	if _, ok := namedLanguage(core.Record{Kind: core.KindValue, Codec: core.CodecUTF8LanguageTextEven}, langs); ok {
		t.Fatal("malformed language text: want false")
	}
}
