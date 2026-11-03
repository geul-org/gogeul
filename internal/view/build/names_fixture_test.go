//ff:func feature=build type=helper control=sequence
//ff:what namesFixture — test fixture: a tree with the Issuer list naming 0x00000000 "geul.org", document 0xffff000000000002 and a codec folder 0x00000000 (an Issuer-looking ID), no document names
package build

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
	doc := filepath.Join(top, "0x00000000", "docs", "0xFFFF000000000002")
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
	os.MkdirAll(filepath.Join(top, "0x00000000", "codecs", "0x00000000"), 0o755)
	return root
}
