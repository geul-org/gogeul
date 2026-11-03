//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestDirPublishAt — versions at an address with any document ID; a republish returns the entry, the next chains; another codec at the codec ID, exhausted, broken-list, unwritable and a dropped LOCAL ID used again fail
package resolve

import (
	"github.com/geul-org/gogeul/pkg/core"
	"math"
	"testing"
)

func TestDirPublishAt(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	k := codecKey(0x0B80)
	doc := codecDocID(0x0B80)
	b := codecDocOf(0x0B80, 4)
	ws, _ := core.ReadWords(b)
	h, _ := core.ContentHash(ws)
	if e, err := d.publishAt(k, []core.Word{0, 0}, doc, b, h); err != nil || e.Version != 1 {
		t.Fatalf("first: %+v %v", e, err)
	}
	if e, err := d.publishAt(k, []core.Word{0, 0}, doc, b, h); err != nil || e.Version != 1 {
		t.Fatalf("again: %+v %v", e, err)
	}
	if lg, err := d.readList(k, core.FormatID(doc)); err != nil || len(lg.Entries) != 1 {
		t.Fatalf("list: %+v %v", lg, err)
	}
	b2 := codecDocOf(0x0B80, 4, 2)
	ws2, _ := core.ReadWords(b2)
	h2, _ := core.ContentHash(ws2)
	if e, err := d.publishAt(k, []core.Word{0, 0}, doc, b2, h2); err != nil || e.Version != 2 {
		t.Fatalf("second: %+v %v", e, err)
	}
	b3 := codecDocOf(0x07A0, 4)
	ws3, _ := core.ReadWords(b3)
	h3, _ := core.ContentHash(ws3)
	if _, err := d.publishAt(k, []core.Word{0, 0}, doc, b3, h3); err == nil {
		t.Fatal("another codec at the same codec ID: want refusal (§1.1 103)")
	}
	ek, eDoc := sampleKeyOf(7), []core.Word{0xffff, 0, 0, 7}
	for i, v := range [][]byte{edgeDoc(7, "a", "b"), edgeDoc(7, "a"), edgeDoc(7, "a", "B")} {
		ws, _ := core.ReadWords(v)
		vh, _ := core.ContentHash(ws)
		if _, err := d.publishAt(ek, []core.Word{0, 0}, eDoc, v, vh); (err != nil) != (i == 2) {
			t.Fatalf("version %d (#2 dropped, then used again): %v", i+1, err)
		}
	}
	full := codecKey(0x0BA0)
	d.write(versionPath(full, math.MaxUint32), b)
	d.write(listPath(full), encodeVersionList([]core.Record{core.EncodePublication(core.Publication{Codec: core.EdgePublication, EdgeID: 1, Issuer: []core.Word{0, 0}, Doc: doc, Version: math.MaxUint32})}, nil, 0, 0))
	if _, err := d.publishAt(full, []core.Word{0, 0}, doc, b, h); err == nil {
		t.Fatal("exhausted: want error")
	}
	junk := codecKey(0x0BB0)
	d.write(listPath(junk), []byte("x"))
	if _, err := d.publishAt(junk, []core.Word{0, 0}, doc, b, h); err == nil {
		t.Fatal("bad list: want error")
	}
	blockWrite(t, d, listPath(codecKey(0x0B90)))
	if _, err := d.publishAt(codecKey(0x0B90), []core.Word{0, 0}, doc, b, h); err == nil {
		t.Fatal("blocked: want error")
	}
}
