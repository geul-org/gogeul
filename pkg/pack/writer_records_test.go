//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterRecords — shared h/ puts are left out; the rest come in record order
package pack

import "testing"

func TestWriterRecords(t *testing.T) {
	p := &put{ref: pathRef{space: spaceP, issuer: "0x00000000"}}
	n := &put{path: "a"}
	h := &put{ref: pathRef{space: spaceH}, via: p}
	w := &writer{inum: map[string]uint32{"0x00000000": 0}, puts: []*put{n, h, p}}
	rs := w.records()
	if len(rs) != 2 || rs[0] != p || rs[1] != n {
		t.Fatal(rs)
	}
}
