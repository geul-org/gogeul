//ff:func feature=notation type=decoder control=sequence
//ff:what TestFormatLanguageText — "text"@tag for a BCP 47 language, "text"@@hex for a short code; a broken record and another language record are false
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestFormatLanguageText(t *testing.T) {
	short := core.Record{Kind: core.KindValue, Codec: core.CodecLanguageShortCode, Payload: []core.Word{0x0038}}
	if s, ok := formatLanguageText(core.EncodeLanguageText(short, []byte("x"))); !ok || s != `"x"@@0x0038` {
		t.Fatalf("short code: %s %v", s, ok)
	}
	bcp := core.EncodeBytesValue(core.CodecLanguageBCP47Even, []byte("ko-KR"))
	if s, ok := formatLanguageText(core.EncodeLanguageText(bcp, []byte("서울"))); !ok || s != `"서울"@ko-KR` {
		t.Fatalf("bcp47: %s %v", s, ok)
	}
	if _, ok := formatLanguageText(core.Record{Kind: core.KindValue, Codec: core.CodecUTF8LanguageTextEven}); ok {
		t.Fatal("broken: want false")
	}
	other := core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("en"))
	if _, ok := formatLanguageText(core.EncodeLanguageText(other, []byte("x"))); ok {
		t.Fatal("UTF8-TEXT as the language: want false")
	}
}
