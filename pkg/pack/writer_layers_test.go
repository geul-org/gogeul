//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterLayers — a delta write keeps older H layers under a new one; P, H and closing failures are returned
package pack

import (
	"errors"
	"testing"
)

func TestWriterLayers(t *testing.T) {
	_, out := built(t, small())
	doc := fixtureFile(t, "0x00000000/docs/0x0000000000000845/0x00000001.geul")
	h, _ := contentHash(doc)
	get := memGet(map[string][]byte{"0x00000000/docs/0x0000000000000001/0x00000001.geul": doc, hPath("0x00000000", string(h[:])): doc})
	h2, _ := contentHash(fixtureFile(t, "0x00000000/docs/0x0000000000000860/0x00000001.geul"))
	mk := func(dels ...string) *writer {
		w := testWriter(t, out, Batch{Put: []string{"0x00000000/docs/0x0000000000000001/0x00000001.geul"}, Delete: append(dels, hPath("0x00000000", string(h[:]))), Get: get})
		w.plan()
		w.writeData()
		return w
	}
	w := mk()
	pl, hl, ns, err := w.layers()
	if err != nil || len(pl) != 1 || len(hl) != 2 || len(ns) != 2 {
		t.Fatal(err, len(hl))
	}
	w = mk()
	w.o.Codec = fakeCodec{err: errors.New("x")}
	if _, _, _, err := w.layers(); err == nil {
		t.Fatal("P")
	}
	w = mk()
	w.o.PLeaf = 1
	w.fs = &fsFuncs{mkdir: -1, create: -1, rename: 0}
	if _, _, _, err := w.layers(); err == nil {
		t.Fatal("close")
	}
	w = mk(hPath("0x00000000", string(h2[:])))
	w.o.PLeaf, w.o.HLeaf = 100, 1
	w.fs = &fsFuncs{create: 0}
	if _, _, _, err := w.layers(); err == nil {
		t.Fatal("H")
	}
}
