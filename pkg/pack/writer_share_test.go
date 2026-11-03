//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterShare — an h/ copy rides on the P record with the same bytes; different bytes or a wrong name stand alone or fail
package pack

import (
	"errors"
	"testing"
)

func TestWriterShare(t *testing.T) {
	doc := fixtureFile(t, "0x00000000/docs/0x0000000000000845/0x00000001.geul")
	h, _ := contentHash(doc)
	hp := hPath("0x00000000", string(h[:]))
	p := &put{path: "0x00000000/docs/0x0000000000000845/0x00000001.geul", ref: parsePath("0x00000000/docs/0x0000000000000845/0x00000001.geul"), sum: [32]byte{1}, ch: h, chOK: true}
	hc := &put{path: hp, ref: parsePath(hp), sum: [32]byte{1}, ch: h, chOK: true}
	w := &writer{puts: []*put{p, hc}}
	if err := w.share(); err != nil || p.share != hp || hc.via != p || w.res.Shared != 1 {
		t.Fatal(err)
	}
	alone := &put{path: hp, ref: parsePath(hp), sum: [32]byte{2}, ch: h, chOK: true}
	p2 := &put{ref: p.ref, sum: [32]byte{1}, ch: h, chOK: true}
	w = &writer{puts: []*put{p2, alone}}
	if err := w.share(); err != nil || alone.via != nil {
		t.Fatal(err)
	}
	wrong := &put{path: hp, ref: parsePath(hp), ch: [32]byte{9}, chOK: true}
	if err := (&writer{puts: []*put{wrong}}).share(); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
