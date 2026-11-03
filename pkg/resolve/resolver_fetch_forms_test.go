//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestResolverFetchForms — Resolver.Fetch binds the same version and bytes from every storage form and passes a missing document to the next source
package resolve

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestResolverFetchForms(t *testing.T) {
	ctx := context.Background()
	tree := formsTree(t)
	seven := []core.Word{0xffff, 0, 0, 7}
	want, err := tree.At([]core.Word{0, 0}, seven, 2)
	if err != nil {
		t.Fatal(err)
	}
	last := Entry{Name: "last", Source: tree}
	for _, f := range formsOf(t, tree) {
		r := &Resolver{Config: Config{Proxy: []Entry{{Name: f.name, Source: f.dir}, last}}}
		n, err := r.Fetch(ctx, Key{Issuer: []core.Word{0, 0}, Doc: seven})
		if err != nil || n.Version != 2 || n.Source != f.name {
			t.Fatalf("%s: %+v %v", f.name, n, err)
		}
		if b, err := f.dir.At([]core.Word{0, 0}, seven, 2); err != nil || !bytes.Equal(b, want) {
			t.Fatalf("%s: At: %v", f.name, err)
		}
		k := Key{Issuer: []core.Word{0, 0}, Doc: []core.Word{0xffff, 0, 0, 9}}
		if _, err := r.Fetch(ctx, k); !errors.Is(err, ErrNotFound) {
			t.Fatalf("%s: missing: %v", f.name, err)
		}
	}
}
