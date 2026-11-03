//ff:func feature=core type=helper control=sequence
//ff:what TestToulminName — the 39 TOULMIN codes have names, SUPERSEDE-CHAIN·VERDICT·EVIDENCE included; Reserved and core codes do not
package core

import "testing"

func TestToulminName(t *testing.T) {
	if len(toulminNames) != 39 {
		t.Fatalf("%d names", len(toulminNames))
	}
	if ToulminName(0x6A) != "SUPERSEDE-CHAIN" || ToulminName(0x78) != "VERDICT" || ToulminName(0x79) != "EVIDENCE" || ToulminName(0x94) != "LABEL" {
		t.Fatal("names")
	}
	if ToulminName(0x7C) != "" || ToulminName(0x95) != "" || ToulminName(EdgeTriple) != "" {
		t.Fatal("not TOULMIN")
	}
}
