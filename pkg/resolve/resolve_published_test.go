//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestResolvePublished — a root importing codec documents of the published geul.org set (0x00000000) resolves through GEULPROXY=file://…,off alone; skipped until that set exists
package resolve

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

// TestResolvePublished reads the codec documents published at geul-dev
// draft/geul-org/artifacts/geul/.well-known/geul/0x00000000/ (the geul.org
// set, the stand-in for geul.org/.well-known/geul/): a root document imports
// the PUBLICATION codec (latest) and pins the FIELD codec at version 1, both by
// CODEC-POINTER. It is skipped while that set has no `latest.txt` (migration.md
// §6.2: the set appears at §7.1 step 5) or when this module is checked out
// without geul-dev draft/ around it. Only those two codecs are read.
func TestResolvePublished(t *testing.T) {
	ctx := context.Background()
	root, _ := filepath.Abs("../../../geul-org/artifacts/geul")
	if _, err := os.Stat(filepath.Join(root, ".well-known", "geul", "0x00000000", "latest.txt")); err != nil {
		t.Skip("geul-org/artifacts/geul/.well-known/geul/0x00000000 holds no pack set yet")
	}
	c, err := ConfigFromEnv(func(k string) string {
		return map[string]string{"GEULPROXY": "file://" + root + ",off"}[k]
	})
	if err != nil {
		t.Fatal(err)
	}
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	id, _ := core.EncodeDocID([]core.Word{0, 0}, []core.Word{0xffff, 0, 0, 7})
	in := enc(id)
	for i, imp := range [][2]uint32{{0x0980, 0}, {0x09C0, 1}} {
		ptr, _ := core.EncodeCodecPointer([]core.Word{0, 0}, imp[0], imp[1])
		in = append(in, enc(core.Record{Kind: core.KindMeta, Codec: core.MetaImportDoc, Payload: append(enc(ptr), core.Word(i+1))})...)
	}
	in = append(append(in, core.TargetWord(core.KindMeta, core.MetaBody)), enc(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("root")))...)
	g, err := (&Resolver{Config: c}).Resolve(ctx, core.PutWords(core.WrapDocument(append(in, core.TargetWord(core.KindMeta, core.MetaEnd)))))
	if err != nil {
		t.Fatal(err)
	}
	key := func(codec core.Word) NodeKey {
		return NodeKey{DocKey: DocKey{Issuer: "0x00000000", Doc: core.FormatID([]core.Word{0, codec})}, Version: 1}
	}
	pub, field := g.Docs[key(0x0980)], g.Docs[key(0x09C0)]
	if len(g.Docs) != 2 || pub == nil || field == nil || pub.Source != "file://"+root {
		t.Fatalf("got %+v", g.Docs)
	}
}
