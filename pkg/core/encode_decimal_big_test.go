//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeDecimalBig — SIGNED-DECIMAL 23.45 and UNSIGNED-DECIMAL 1256.1 give the grammar's words; a negative unsigned coefficient fails
package core

import (
	"math/big"
	"testing"
)

func TestEncodeDecimalBig(t *testing.T) {
	r, err := EncodeDecimalBig(true, -2, big.NewInt(2345))
	ws, _ := EncodeRecord(r)
	if err != nil || FormatID(ws) != "0x0483fffffffe0929" {
		t.Fatalf("signed: %s %v", FormatID(ws), err)
	}
	r, err = EncodeDecimalBig(false, -1, big.NewInt(12561))
	ws, _ = EncodeRecord(r)
	if err != nil || FormatID(ws) != "0x0493ffffffff3111" {
		t.Fatalf("unsigned: %s %v", FormatID(ws), err)
	}
	if _, err := EncodeDecimalBig(false, 0, big.NewInt(-1)); err == nil {
		t.Fatal("negative unsigned: want error")
	}
}
