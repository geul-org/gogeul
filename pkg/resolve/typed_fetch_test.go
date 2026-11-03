//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestTypedFetch — Issuers, Docs, Doc, DocVersions, Codec, CodecVersions, Codecs, CodecsVersions, IssuersVersions, DocsVersions and Hash read their addresses; a padded Issuer is the same Issuer
package resolve

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestTypedFetch(t *testing.T) {
	ctx := context.Background()
	d := Dir{Root: t.TempDir()}
	d.Publish(edgeDoc(7, "a"))
	d.PublishCodec(codecDocOf(0x0980, 4))
	d.Name([]core.Word{0, 0}, []core.Word{0xffff, 0, 0, 7}, "seven")
	issuers, _ := namesDoc(nil, nil, core.Bind{EdgeID: 1, Subject: func() core.Record { r, _ := core.EncodeIssuerPointer([]core.Word{0, 0}); return r }(), Name: []byte("geul.org")})
	d.write(issuersPath(), issuers)
	g, pad := []core.Word{0, 0}, []core.Word{0, 0, 0, 0}
	v7 := edgeDoc(7, "a")
	ws, _ := core.ReadWords(v7)
	h, _ := core.ContentHash(ws)
	list, _ := d.Fetch(ctx, listPath(sampleKeyOf(7)))
	clist, _ := d.Fetch(ctx, listPath(codecKey(0x0980)))
	names, _ := d.Fetch(ctx, namesPath(geulIssuer))
	ok := func(name string, b []byte, err error, want []byte) {
		if err != nil || !bytes.Equal(b, want) {
			t.Fatalf("%s: %d bytes %v", name, len(b), err)
		}
	}
	b, err := Issuers(ctx, d, Latest)
	ok("Issuers", b, err, issuers)
	b, err = Docs(ctx, d, pad, 1)
	ok("Docs", b, err, names)
	b, err = Doc(ctx, d, pad, 0xffff000000000007, Latest)
	ok("Doc", b, err, v7)
	b, err = DocVersions(ctx, d, g, 0xffff000000000007)
	ok("DocVersions", b, err, list)
	b, err = Codec(ctx, d, g, 0x0980, 1)
	ok("Codec", b, err, codecDocOf(0x0980, 4))
	b, err = CodecVersions(ctx, d, g, 0x0980)
	ok("CodecVersions", b, err, clist)
	b, err = Hash(ctx, d, g, h)
	ok("Hash", b, err, v7)
	for name, f := range map[string]func() ([]byte, error){
		"IssuersVersions": func() ([]byte, error) { return IssuersVersions(ctx, d) },
		"CodecsVersions":  func() ([]byte, error) { return CodecsVersions(ctx, d, g) },
		"DocsVersions":    func() ([]byte, error) { return DocsVersions(ctx, d, g) },
		"Codecs":          func() ([]byte, error) { return Codecs(ctx, d, g, Latest) },
		"Hash of another": func() ([]byte, error) { return Hash(ctx, d, []core.Word{0, 1}, h) },
		"1-word Issuer":   func() ([]byte, error) { return Doc(ctx, d, []core.Word{0}, 0xffff000000000007, Latest) },
	} {
		if _, err := f(); !errors.Is(err, ErrNotFound) {
			t.Fatalf("%s: %v", name, err)
		}
	}
}
