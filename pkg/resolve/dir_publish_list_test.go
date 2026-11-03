//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestDirPublishList — list versions from 1 with versions.geul pointing DOC-POINTER(Issuer, reserved ID, version); same bytes again is that version; non-list IDs, a non-geul.org Issuer list and non-documents fail
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
	is := []core.Word{0, 2}
	ip, _ := core.EncodeIssuerPointer(is)
	b1, _ := namesDoc(nil, nil, core.Bind{EdgeID: 1, Subject: ip, Name: []byte("a")})
	b2, _ := namesDoc(nil, nil, core.Bind{EdgeID: 1, Subject: ip, Name: []byte("b")})
	for i, b := range [][]byte{b1, b2, b2} {
		e, err := d.PublishList(is, core.DocIDDocs, b)
		if err != nil || e.Version != uint32(min(i+1, 2)) {
			t.Fatal(i, e, err)
		}
	}
	got, err := Docs(context.Background(), d, is, Latest)
	if err != nil || !bytes.Equal(got, b2) {
		t.Fatal(len(got), err)
	}
	lb, _ := d.Fetch(context.Background(), "0x00000002/docs/versions.geul")
	ws, _ := core.ReadWords(lb)
	doc, err := core.SplitDocument(ws)
	if err != nil || len(doc.Body) != 2 {
		t.Fatal(err)
	}
	p, err := core.DecodePublication(doc.Body[1])
	if err != nil || !slices.Equal(p.Issuer, is) || core.FormatID(p.Doc) != listIDDocs || p.Version != 2 {
		t.Fatalf("%+v %v", p, err)
	}
	bads := []struct {
		is   []core.Word
		list uint64
		b    []byte
	}{{is, 7, b1}, {is, core.DocIDIssuers, b1}, {is, core.DocIDCodecs, []byte{1}}, {is, core.DocIDCodecs, []byte{0, 1}}}
	for _, c := range bads {
		if _, err := d.PublishList(c.is, c.list, c.b); err == nil {
			t.Fatalf("%x: want error", c.list)
		}
	}
}
