//ff:func feature=schema type=decoder control=sequence
//ff:what TestMessageLanguage — the language record of a MESSAGE; no payload, another Value and a truncated language record give false
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestMessageLanguage(t *testing.T) {
	msg := func(v core.Record) core.Record {
		return core.Record{Kind: core.KindMeta, Codec: core.MetaMessage, Payload: enc(v)}
	}
	if l, ok := messageLanguage(msg(core.Record{Kind: core.KindValue, Codec: core.CodecUTF8LanguageTextEven | 1, Payload: []core.Word{0x07A1, 9, 0x6800}})); !ok || l != "0x07A10009" {
		t.Fatal(l, ok)
	}
	if _, ok := messageLanguage(core.Record{Kind: core.KindMeta, Codec: core.MetaMessage}); ok {
		t.Fatal("empty")
	}
	if _, ok := messageLanguage(msg(core.EncodeUintValue(1))); ok {
		t.Fatal("other value")
	}
	if _, ok := messageLanguage(msg(core.Record{Kind: core.KindValue, Codec: core.CodecUTF8LanguageTextEven, Payload: []core.Word{0x07A3, 9}})); ok {
		t.Fatal("truncated language")
	}
}
