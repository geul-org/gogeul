//ff:func feature=resolve type=resolver control=sequence
//ff:what TestResolverCodecSpecs — a spec missing from the first source comes from the next; a target no source has is left out; a broken source stops
package resolve

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestResolverCodecSpecs(t *testing.T) {
	const tag = core.Word(0x0980)
	published := Dir{Root: t.TempDir()}
	if _, err := published.PublishCodec(codecDocOf(tag, 4, 2)); err != nil {
		t.Fatal(err)
	}
	empty := Dir{Root: t.TempDir()}
	r := Resolver{Config: Config{Proxy: []Entry{{Name: "data", Source: empty}, {Name: "geul.org", Source: published}}}}
	s, err := r.CodecSpecs(t.Context(), []core.Word{tag, 0x0990})
	if err != nil || s[tag] == nil || len(s) != 1 {
		t.Fatalf("%+v %v", s, err)
	}
	r.Config.Proxy = []Entry{{Name: "data", Source: empty}}
	if s, err := r.CodecSpecs(t.Context(), []core.Word{tag}); err != nil || len(s) != 0 {
		t.Fatalf("%+v %v", s, err)
	}
	broken := Dir{Root: t.TempDir()}
	broken.write(listPath(codecKey(tag)), []byte{1})
	r.Config.Proxy = []Entry{{Name: "broken", Source: broken}, {Name: "geul.org", Source: published}}
	if _, err := r.CodecSpecs(t.Context(), []core.Word{tag}); err == nil {
		t.Fatal("broken source: want error")
	}
}
