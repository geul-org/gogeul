//ff:func feature=server type=helper control=sequence
//ff:what namesFixture — test fixture: a tree with the Issuer list naming 0x00000000 "geul.org" and its document list naming document 2 "sample-b" and codec 0x00000980 "publication"
package server

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/resolve"
)

func namesFixture(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	top := filepath.Join(root, ".well-known", "geul")
	doc := filepath.Join(top, "0x00000000", "docs", "0xffff000000000002")
	if err := os.MkdirAll(doc, 0o755); err != nil {
		t.Fatal(err)
	}
	os.MkdirAll(filepath.Join(top, "issuers"), 0o755)
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	subj := core.Record{Kind: core.KindReference, Codec: core.CodecIssuerPointer, Payload: []core.Word{0, 0}}
	p := append(enc(core.EncodeLocalRef(1)), enc(subj)...)
	p = append(p, enc(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("geul.org")))...)
	in := append([]core.Word{core.TargetWord(core.KindMeta, core.MetaBody)}, enc(core.Record{Kind: core.KindEdge, Codec: core.EdgeBind, Payload: p})...)
	end := core.TargetWord(core.KindMeta, core.MetaEnd)
	d := resolve.Dir{Root: root}
	if _, err := d.PublishList([]core.Word{0, 0}, core.DocIDIssuers, core.PutWords(core.WrapDocument(append(in, end)))); err != nil {
		t.Fatal(err)
	}
	os.WriteFile(filepath.Join(doc, "0x00000001.geul"), core.PutWords(core.WrapDocument([]core.Word{end})), 0o644)
	dsubj := core.Record{Kind: core.KindReference, Codec: core.CodecDocPointer, Payload: []core.Word{0, 0, 0xffff, 0, 0, 2, 0, 0}}
	dp := append(enc(core.EncodeLocalRef(1)), enc(dsubj)...)
	dp = append(dp, enc(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("sample-b")))...)
	csubj := core.Record{Kind: core.KindReference, Codec: core.CodecCodecPointer, Payload: []core.Word{0, 0, 0, 0x0980, 0, 0}}
	cp := append(enc(core.EncodeLocalRef(2)), enc(csubj)...)
	cp = append(cp, enc(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("publication")))...)
	din := append([]core.Word{core.TargetWord(core.KindMeta, core.MetaBody)}, enc(core.Record{Kind: core.KindEdge, Codec: core.EdgeBind, Payload: dp})...)
	din = append(din, enc(core.Record{Kind: core.KindEdge, Codec: core.EdgeBind, Payload: cp})...)
	if _, err := d.PublishList([]core.Word{0, 0}, core.DocIDDocs, core.PutWords(core.WrapDocument(append(din, end)))); err != nil {
		t.Fatal(err)
	}
	os.MkdirAll(filepath.Join(top, "0x00000000", "codecs", "0x00000980"), 0o755)
	return root
}
