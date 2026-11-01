//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestReadHeads — heads come back from their place in the file; reading past its end is an error
package pack

import (
	"bytes"
	"testing"
)

func TestReadHeads(t *testing.T) {
	var b []byte
	for i := 0; i < 3; i++ {
		var e hEntry
		e.key[0], e.loc.Off = byte(i), uint64(i)
		b = appendHEntry(b, e)
	}
	got, err := readHeads(bytes.NewReader(b), 1, 2)
	if err != nil || len(got) != 2 || got[0].Key[0] != 1 || got[1].Loc.Off != 2 || len(got[0].Key) != 32 {
		t.Fatalf("%v %v", got, err)
	}
	if _, err := readHeads(bytes.NewReader(b), 2, 2); err == nil {
		t.Fatal("past the end: want error")
	}
}
