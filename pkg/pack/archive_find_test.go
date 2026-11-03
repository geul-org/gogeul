//ff:func feature=pack type=helper control=sequence
//ff:what TestArchiveFind — by name, nil when missing
package pack

import "testing"

func TestArchiveFind(t *testing.T) {
	a := &Archive{sets: []*packSet{{name: "0x00000000"}}}
	if a.find("0x00000000") != a.sets[0] || a.find("") != nil {
		t.Fatal("find")
	}
}
