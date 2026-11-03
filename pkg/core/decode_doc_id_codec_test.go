//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodeDocIDCodec — the DOC-ID 0x1C77 0x0426 … of a codec document gives Issuer and codec ID; a DOC-POINTER DOC-ID, a version, a broken or longer payload are rejected
package core

import (
	"slices"
	"testing"
)

func TestDecodeDocIDCodec(t *testing.T) {
	r, _, err := DecodeRecord([]Word{0x1C77, 0x0426, 0, 0, 0, 0x0B80, 0, 0})
	if err != nil {
		t.Fatal(err)
	}
	if iss, c, err := DecodeDocIDCodec(r); err != nil || !slices.Equal(iss, []Word{0, 0}) || c != 0x0B80 {
		t.Fatalf("%04X %x %v", iss, c, err)
	}
	bads := [][]Word{
		{0x1C79, 0x0418, 0, 0, 0, 0, 0, 1, 0, 0},
		{0x1C77, 0x0426, 0, 0, 0, 0x0B80, 0, 1},
		{0x1C77, 0x0426, 0, 0, 0, 0x0B80, 0x0008, 1},
		{0x1C71, 0x0008},
		{0x0601, 1},
	}
	for _, b := range bads {
		r, _, err := DecodeRecord(b)
		if err != nil {
			t.Fatal(err)
		}
		if _, _, err := DecodeDocIDCodec(r); err == nil {
			t.Fatalf("%04X accepted", b)
		}
	}
}
