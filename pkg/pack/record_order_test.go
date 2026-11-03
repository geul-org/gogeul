//ff:func feature=pack type=writer control=sequence
//ff:what TestRecordOrder — P records by key first, then h/-only by hash, then named by path
package pack

import "testing"

func TestRecordOrder(t *testing.T) {
	inum := map[string]uint32{"0x00000000": 0}
	p := recordOrder(&put{ref: parsePath("0x00000000/docs/0x0000000000000008/versions.geul")}, inum)
	h := recordOrder(&put{ref: pathRef{space: spaceH, hash: "zz"}}, inum)
	n := recordOrder(&put{path: "a", ref: pathRef{}}, inum)
	if !(p < h && h < n) || n != "\x03a" {
		t.Fatal(p, h, n)
	}
}
