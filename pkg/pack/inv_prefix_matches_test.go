//ff:func feature=pack type=writer control=sequence
//ff:what TestInvPrefixMatches — entries that begin with the checkpoint's match field for field; one differing field does not
package pack

import "testing"

func TestInvPrefixMatches(t *testing.T) {
	entries := []invEntry{{kind: invP, enc: encZstd, flags: flagHash, ref: 0, doc: 7, slot: 1, n: 10, off: 64}, {kind: invNamed, off: 74}}
	inv := []streamInv{{Kind: invP, Enc: encZstd, Flags: flagHash, Doc: 7, Slot: 1, N: 10, Off: 64}}
	if !invPrefixMatches(entries, inv) || !invPrefixMatches(entries, nil) {
		t.Fatal("prefix: want match")
	}
	inv[0].Slot = 2
	if invPrefixMatches(entries, inv) {
		t.Fatal("other slot: want no match")
	}
}
