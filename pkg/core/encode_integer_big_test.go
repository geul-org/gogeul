//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeIntegerBig — SIGNED-INTEGER 2 is 0881 0002, UNSIGNED-INTEGER 3 is 0882 0003; a negative unsigned fails
package core

import (
	"math/big"
	"testing"
)

func TestEncodeIntegerBig(t *testing.T) {
	r, err := EncodeIntegerBig(true, big.NewInt(2))
	ws, _ := EncodeRecord(r)
	if err != nil || FormatID(ws) != "0x06110002" {
		t.Fatalf("signed: %s %v", FormatID(ws), err)
	}
	r, err = EncodeIntegerBig(false, big.NewInt(3))
	ws, _ = EncodeRecord(r)
	if err != nil || FormatID(ws) != "0x06210003" {
		t.Fatalf("unsigned: %s %v", FormatID(ws), err)
	}
	if _, err := EncodeIntegerBig(false, big.NewInt(-3)); err == nil {
		t.Fatal("negative unsigned: want error")
	}
}
