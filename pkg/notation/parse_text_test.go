//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what TestParseText — UTF8-TEXT, language text with a BCP 47 tag or a short code (@@hex, @@registry tag); a broken quote, a bare @, junk after the quote and an unregistered @@tag fail
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
)

func TestParseText(t *testing.T) {
	if r, err := parseText(`"a\"b"`, Context{}); err != nil || r.Codec&^1 != core.CodecUTF8TextEven {
		t.Fatalf("%+v %v", r, err)
	}
	r, err := parseText(`"Adams"@en`, Context{})
	if s, ok := formatLanguageText(r); err != nil || !ok || s != `"Adams"@en` {
		t.Fatalf("%s %v", s, err)
	}
	r, err = parseText(`"Adams"@@0x0784`, Context{})
	if s, ok := formatLanguageText(r); err != nil || !ok || s != `"Adams"@@0x0784` {
		t.Fatalf("hex: %s %v", s, err)
	}
	ctx := Context{Spec: codecspec.Spec{0x05A0: {{Name: "code", Width: 1, Enums: map[string]string{"0x4130": "pt-BR"}}}}}
	r, err = parseText(`"Brasil"@@pt-br`, ctx)
	if s, ok := formatLanguageText(r); err != nil || !ok || s != `"Brasil"@@0x4130` {
		t.Fatalf("registry tag: %s %v", s, err)
	}
	for _, s := range []string{`"a`, `"a"@`, `"a"x`, `"a"@@de-x-formal`, `"a"@@`} {
		if _, err := parseText(s, ctx); err == nil {
			t.Errorf("%s: want error", s)
		}
	}
}
