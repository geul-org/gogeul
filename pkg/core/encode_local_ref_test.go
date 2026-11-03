//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeLocalRef — TAG WORD `0000 0LLL L100 0000` and minimal ID words
package core

import "testing"

func TestEncodeLocalRef(t *testing.T) {
	ws, err := EncodeRecord(EncodeLocalRef(3))
	if err != nil || len(ws) != 2 || ws[0] != 0x0051 || ws[1] != 3 {
		t.Fatalf("got %04X %v", ws, err)
	}
	if r := EncodeLocalRef(70000); len(r.Payload) != 2 {
		t.Fatalf("%+v", r)
	}
}
