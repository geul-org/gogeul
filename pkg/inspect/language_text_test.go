//ff:func feature=inspect type=decoder control=sequence
//ff:what TestLanguageText — BCP 47 and short-code languages decode with their text; another record and a broken value fail
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestLanguageText(t *testing.T) {
	fr := core.EncodeBytesValue(core.CodecLanguageBCP47Even, []byte("fr"))
	if text, lang, err := languageText(core.EncodeLanguageText(fr, []byte("Paris"))); err != nil || text != "Paris" || lang != "fr" {
		t.Fatalf("%q %q %v", text, lang, err)
	}
	short := core.Record{Kind: core.KindValue, Codec: core.CodecLanguageShortCode, Payload: []core.Word{0x0783}}
	if text, lang, err := languageText(core.EncodeLanguageText(short, []byte("DNA"))); err != nil || text != "DNA" || lang != "0x0783" {
		t.Fatalf("short: %q %q %v", text, lang, err)
	}
	if _, _, err := languageText(core.EncodeLanguageText(core.EncodeUintValue(1), []byte("x"))); err == nil {
		t.Fatal("not a language: want error")
	}
	if _, _, err := languageText(core.EncodeUintValue(1)); err == nil {
		t.Fatal("not language text: want error")
	}
}
