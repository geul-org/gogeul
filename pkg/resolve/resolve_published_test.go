//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestResolvePublished — a root importing codec spec documents of the published pack archive (geul-org/artifacts/geul) resolves through GEULPROXY=file://…,off alone, nothing injected
package resolve

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

// TestResolvePublished reads the codec spec documents published at geul-dev
// draft/geul-org/artifacts/geul/.well-known/geul/ (pack sets, the stand-in for
// geul.org/.well-known/geul/): a root document imports the PUBLICATION spec
// document (latest) and pins the FIELD one at version 1. It is skipped when
// this module is checked out on its own, without geul-dev draft/ around it.
func TestResolvePublished(t *testing.T) {
	ctx := context.Background()
	root, _ := filepath.Abs("../../../geul-org/artifacts/geul")
	if _, err := os.Stat(filepath.Join(root, ".well-known", "geul", "0x0000", "latest.txt")); err != nil {
		t.Skip("geul-org/artifacts/geul/.well-known/geul holds no pack sets")
	}
	c, err := ConfigFromEnv(func(k string) string {
		return map[string]string{"GEULPROXY": "file://" + root + ",off"}[k]
	})
	if err != nil {
		t.Fatal(err)
	}
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	id, _ := core.EncodeDocID([]core.Word{0}, []core.Word{0xffff, 0, 0, 7})
	in := enc(id)
	for i, imp := range [][2]core.Word{{0x1018, 0}, {0x101C, 1}} {
		ptr, _ := core.EncodeDocPointer([]core.Word{0}, []core.Word{0, 0, 0, imp[0]}, []core.Word{0, imp[1]})
		in = append(in, enc(core.Record{Kind: core.KindMeta, Codec: core.MetaImportDoc, Payload: append(enc(ptr), core.Word(i+1))})...)
	}
	in = append(append(in, 0x180E), enc(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("root")))...)
	g, err := (&Resolver{Config: c}).Resolve(ctx, core.PutWords(core.WrapDocument(append(in, 0x180F))))
	if err != nil {
		t.Fatal(err)
	}
	key := func(w core.Word) NodeKey {
		return NodeKey{DocKey: DocKey{Issuer: core.FormatID([]core.Word{0}), Doc: core.FormatID([]core.Word{0, 0, 0, w})}, Version: 1}
	}
	pub, field := g.Docs[key(0x1018)], g.Docs[key(0x101C)]
	if len(g.Docs) != 2 || pub == nil || field == nil || pub.Source != "file://"+root {
		t.Fatalf("got %+v", g.Docs)
	}
}
