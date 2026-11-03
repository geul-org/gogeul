//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeFraction — 2/3 and 23.45/1256.1 give the grammar's words; a record that does not encode fails
package core

import (
	"math/big"
	"testing"
)

func TestEncodeFraction(t *testing.T) {
	n, _ := EncodeIntegerBig(true, big.NewInt(2))
	d, _ := EncodeIntegerBig(false, big.NewInt(3))
	r, err := EncodeFraction(Fraction{Numerator: n, Denominator: d})
	ws, _ := EncodeRecord(r)
	if err != nil || FormatID(ws) != "0x04740411000204210003" {
		t.Fatalf("2/3: %s %v", FormatID(ws), err)
	}
	sd, _ := EncodeDecimalBig(true, -2, big.NewInt(2345))
	ud, _ := EncodeDecimalBig(false, -1, big.NewInt(12561))
	r, _ = EncodeFraction(Fraction{Numerator: sd, Denominator: ud})
	ws, _ = EncodeRecord(r)
	if FormatID(ws) != "0x04780483fffffffe09290493ffffffff3111" {
		t.Fatalf("23.45/1256.1: %s", FormatID(ws))
	}
	if _, err := EncodeFraction(Fraction{Numerator: Record{Kind: KindValue, Codec: CodecValueIssuer2Word}, Denominator: d}); err == nil {
		t.Fatal("bad record: want error")
	}
}
