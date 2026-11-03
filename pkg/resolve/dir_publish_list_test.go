//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestDirPublishList — Issuer list versions from 1 with versions.geul pointing DOC-POINTER(0x00000000, reserved ID, version); same bytes again is that version; the gone docs/codecs list IDs, other IDs, another Issuer and non-documents fail
package resolve

import (
	"bytes"
	"context"
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDirPublishList(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	g := []core.Word{0, 0}
	ip, _ := core.EncodeIssuerPointer([]core.Word{0, 2})
	b1 := bindDoc(core.Bind{Subject: ip, Name: []byte("a")})
	b2 := bindDoc(core.Bind{Subject: ip, Name: []byte("b")})
	for i, b := range [][]byte{b1, b2, b2} {
		e, err := d.PublishList(g, core.DocIDIssuers, b)
		if err != nil || e.Version != uint32(min(i+1, 2)) {
			t.Fatal(i, e, err)
		}
	}
	got, err := Issuers(context.Background(), d, Latest)
	if err != nil || !bytes.Equal(got, b2) {
		t.Fatal(len(got), err)
	}
	lb, _ := d.Fetch(context.Background(), "issuers/versions.geul")
	ws, _ := core.ReadWords(lb)
	doc, err := core.SplitDocument(ws)
	if err != nil || len(doc.Body) != 2 {
		t.Fatal(err)
	}
	p, err := core.DecodePublication(doc.Body[1])
	if err != nil || !slices.Equal(p.Issuer, g) || core.FormatID(p.Doc) != listIDIssuers || p.Version != 2 {
		t.Fatalf("%+v %v", p, err)
	}
	bads := []struct {
		is   []core.Word
		list uint64
		b    []byte
	}{{g, 7, b1}, {g, 0xFFFFFFFFFFFFFFFE, b1}, {g, 0xFFFFFFFFFFFFFFFD, b1}, {[]core.Word{0, 2}, core.DocIDIssuers, b1},
		{g, core.DocIDIssuers, []byte{1}}, {g, core.DocIDIssuers, []byte{0, 1}}}
	for _, c := range bads {
		if _, err := d.PublishList(c.is, c.list, c.b); err == nil {
			t.Fatalf("%x: want error", c.list)
		}
	}
}
