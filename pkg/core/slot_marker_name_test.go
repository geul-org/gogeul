//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestSlotMarkerName — C 0/1/2 are ABSENT/UNKNOWN-VALUE/NO-VALUE; pointers and LOCAL are not markers
package core

import "testing"

func TestSlotMarkerName(t *testing.T) {
	want := []string{"ABSENT", "UNKNOWN-VALUE", "NO-VALUE", "", "", "", "", "", "", ""}
	for c, name := range want {
		if got := SlotMarkerName(uint16(c)); got != name {
			t.Fatalf("SlotMarkerName(%d)=%q want %q", c, got, name)
		}
	}
}
