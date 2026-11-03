//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterEntryOf — kind 1 with Issuer, document, slot and share flag; kind 2; kind 3 with a path table entry
package pack

import "testing"

func TestWriterEntryOf(t *testing.T) {
	w := &writer{inum: map[string]uint32{"0x00000000": 0, "0x00000002": 1}}
	e := w.entryOf(&put{ref: parsePath("0x00000002/docs/0x0000000000000008/0x00000002.geul"), share: "h", loc: Loc{Off: 70, Len: 3}})
	if e.kind != invP || e.ref != 1 || e.doc != 8 || e.slot != 2 || e.flags != flagHash || e.off != 70 || e.n != 3 {
		t.Fatal(e)
	}
	if e := w.entryOf(&put{ref: pathRef{space: spaceH}}); e.kind != invH || e.flags != 0 {
		t.Fatal(e)
	}
	if e := w.entryOf(&put{path: "a.geul"}); e.kind != invNamed || e.ref != 0 || w.inv.paths[0] != "a.geul" {
		t.Fatal(e)
	}
}
