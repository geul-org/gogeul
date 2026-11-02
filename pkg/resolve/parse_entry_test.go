//ff:func feature=resolve type=decoder control=iteration dimension=1
//ff:what TestParseEntry — off, registry, file:// and https:// as sites whose form is found when read (geul.org public, others not); http:// and junk fail
package resolve

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseEntry(t *testing.T) {
	for s, want := range map[string]Entry{
		"off":      {Name: "off", Off: true},
		"registry": {Name: "registry", Source: Registry{}, Public: true},
	} {
		if e, err := parseEntry(s); err != nil || !reflect.DeepEqual(e, want) {
			t.Fatalf("%q: %+v %v", s, e, err)
		}
	}
	f, err := parseEntry("file:///srv")
	d, ok := f.Source.(Dir)
	if s, isSite := d.Store.(*site); err != nil || !ok || !isSite || d.Root != "/srv" || f.Name != "file:///srv" || f.Public || s.dir != "/srv/.well-known/geul" {
		t.Fatalf("file: %+v %v", f, err)
	}
	for s, public := range map[string]bool{"https://geul.org/": true, "https://geul.example.com": false} {
		e, err := parseEntry(s)
		d, ok := e.Source.(Dir)
		st, isSite := d.Store.(*site)
		if err != nil || !ok || !isSite || d.Root != "" || st.dir != "" || e.Public != public || e.Name != s {
			t.Fatalf("%q: %+v %v", s, e, err)
		}
		if h, ok := st.loose.(HTTP); !ok || h.Base != strings.TrimSuffix(s, "/") {
			t.Fatalf("%q: loose %+v", s, st.loose)
		}
	}
	for _, s := range []string{"http://x", "file://", "", "geul.org"} {
		if _, err := parseEntry(s); err == nil {
			t.Fatalf("%q: want error", s)
		}
	}
}
