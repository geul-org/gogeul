//ff:func feature=cli type=helper control=sequence
//ff:what TestAppDocSpec — only the codecs the document uses; the language registry though no top-level record has its codec; same lines as all specs; empty without a data folder, on bad bytes and on a broken spec
package cli

import (
	"io"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestAppDocSpec(t *testing.T) {
	h := newHarness(t)
	h.languageSpec()
	addr, draft := h.doc()
	h.must("add", draft, "ALIAS", "Q42", `"Doug"@@en`)
	b, err := os.ReadFile(draft)
	if err != nil {
		t.Fatal(err)
	}
	a := h.app(io.Discard)
	s := a.docSpec(b)
	if len(s) != 2 || s[core.TargetWord(core.KindEdge, 0x0A)] == nil || s[0x0848] == nil {
		t.Fatalf("want TRIPLE and the registry: %+v", s)
	}
	got, err := linesWith(b, s)
	want, _ := linesWith(b, a.spec())
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("%v\n got %q\nwant %q", err, got, want)
	}
	v1, err := h.app(io.Discard).readDoc(addr + "@1")
	if err != nil {
		t.Fatal(err)
	}
	if s := a.docSpec(v1); s[0x0848] != nil {
		t.Fatal("no language code in this version: the registry is not read")
	}
	if len(a.docSpec([]byte{1})) != 0 {
		t.Fatal("unreadable bytes: want empty")
	}
	if len(a.docSpec(make([]byte, 8))) != 0 {
		t.Fatal("not a document: want empty")
	}
	list := filepath.Join(h.data, ".well-known", "geul", "0x0000", "0x000000000000100a", "versions.geul")
	os.WriteFile(list, []byte{1}, 0o644)
	if len(a.docSpec(b)) != 0 {
		t.Fatal("broken spec: want empty")
	}
	h.cwd = t.TempDir()
	if len(h.app(io.Discard).docSpec(b)) != 0 {
		t.Fatal("no data folder: want empty")
	}
}
