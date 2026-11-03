//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what newTypedFixture — publish the typedFixture tree and read back the stored bytes each typed fetch must return
package resolve

import (
	"context"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

// newTypedFixture builds the typedFixture tree under a temporary root.
func newTypedFixture(t *testing.T) typedFixture {
	t.Helper()
	ctx := context.Background()
	f := typedFixture{d: Dir{Root: t.TempDir()}, g: []core.Word{0, 0}, pad: []core.Word{0, 0, 0, 0}}
	f.v1, f.v2, f.codec = edgeDoc(7, "a"), edgeDoc(7, "b"), codecDocOf(0x0B80, 4)
	for _, b := range [][]byte{f.v1, f.v2} {
		if _, err := f.d.Publish(b); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := f.d.PublishCodec(f.codec); err != nil {
		t.Fatal(err)
	}
	ip, _ := core.EncodeIssuerPointer(f.g)
	f.issuers = bindDoc(core.Bind{Subject: ip, Name: []byte("geul.org")})
	if _, err := f.d.PublishList(f.g, core.DocIDIssuers, f.issuers); err != nil {
		t.Fatal(err)
	}
	ws, _ := core.ReadWords(f.v2)
	f.hash, _ = core.ContentHash(ws)
	read := func(p string) []byte { b, _ := f.d.Fetch(ctx, p); return b }
	f.list, f.clist = read(listPath(sampleKeyOf(7))), read(listPath(codecKey(0x0B80)))
	f.ilist = read(listPath(issuersKey()))
	return f
}
