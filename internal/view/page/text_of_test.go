//ff:func feature=page type=builder control=iteration dimension=1
//ff:what TestTextOf — UTF8-TEXT decodes, UTF8-LANGUAGE-TEXT as "text"@lang; other Values, Kinds and bad padding are not text
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestTextOf(t *testing.T) {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	if s, ok := textOf(enc(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("geul.org")))); !ok || s != "geul.org" {
		t.Fatalf("utf8: %q %v", s, ok)
	}
	fr := core.EncodeBytesValue(core.CodecLanguageBCP47Even, []byte("fr"))
	if s, ok := textOf(enc(core.EncodeLanguageText(fr, []byte("Paris")))); !ok || s != `"Paris"@fr` {
		t.Fatalf("language text: %q %v", s, ok)
	}
	for i, ws := range [][]core.Word{enc(core.EncodeUintValue(1)), enc(core.EncodeLocalRef(1)), {0x0AC1, 0x4142}, {0xE000}, enc(core.EncodeBytesValue(core.CodecURIEven, []byte("http://x"))), {0x0731, 0x4142}} {
		if _, ok := textOf(ws); ok {
			t.Fatalf("case %d: want not text", i)
		}
	}
}
