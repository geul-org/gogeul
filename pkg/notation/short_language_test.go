//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what TestShortLanguage — a registry tag (case ignored) or an assigned 1-WORD hex gives a LANGUAGE-SHORT-CODE; an unregistered tag or code, a tag without registry, empty, 0x0000, 0xffff and 2 WORDs fail; hex without registry is kept
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
)

func TestShortLanguage(t *testing.T) {
	ctx := Context{Spec: codecspec.Spec{0x05A0: {{Name: "code", Width: 1, Enums: map[string]string{"0x404c": "zh-Hans", "0x0784": "en", "0x0038": "ach"}}}}}
	for in, want := range map[string]core.Word{"zh-hans": 0x404C, "EN": 0x0784, "0x0038": 0x0038} {
		r, err := shortLanguage(in, ctx)
		if err != nil || r.Kind != core.KindValue || r.Codec != core.CodecLanguageShortCode || len(r.Payload) != 1 || r.Payload[0] != want {
			t.Errorf("%s: %+v %v", in, r, err)
		}
	}
	for _, in := range []string{"de-x-formal", "", "0x0000", "0xffff", "0x00010002", "0xzz", "0x7fff"} {
		if _, err := shortLanguage(in, ctx); err == nil {
			t.Errorf("%q: want error", in)
		}
	}
	if _, err := shortLanguage("en", Context{}); err == nil {
		t.Error("no registry: want error")
	}
	if r, err := shortLanguage("0x7fff", Context{}); err != nil || r.Payload[0] != 0x7FFF {
		t.Errorf("no registry, hex kept as written: %+v %v", r, err)
	}
}
