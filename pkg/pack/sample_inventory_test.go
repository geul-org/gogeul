//ff:func feature=pack type=helper control=sequence
//ff:what sampleInventory — an inventory with every part filled
package pack

import "strings"

// sampleInventory returns a small complete inventory.
func sampleInventory() inventory {
	return inventory{
		issuers: []string{"0x00000000"},
		paths:   []string{"issuers/0x00000001.geul"},
		entries: []invEntry{{kind: invP, flags: flagHash, doc: 8, slot: 1, n: 3, off: 64}, {kind: invNamed, n: 2, off: 67}, {kind: invPDel, doc: 9}},
		deleted: []string{strings.Repeat("d", 32)},
	}
}
