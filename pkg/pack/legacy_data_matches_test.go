//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestLegacyDataMatches — a sealed pack matches each hash-less checkpoint of it (mid-pack and full); another size, count, entry or path, an empty inventory, a failing or short read, a bad header and a bad inventory frame do not
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLegacyDataMatches(t *testing.T) {
	_, sealed, snaps := streamTestSealed(t)
	f, _ := os.Open(sealed)
	defer f.Close()
	mid, full := snaps[len(snaps)/2], snaps[len(snaps)-1]
	for _, sn := range []streamSnap{mid, full} {
		if !legacyDataMatches(f, sn.open, sn.inv, sn.paths) {
			t.Fatalf("checkpoint at %d: want match", sn.open.Off)
		}
	}
	off := mid.open
	off.Off++
	more := full.open
	more.Count++
	other := append([]streamInv(nil), mid.inv...)
	other[1].Doc++
	for name, ok := range map[string]bool{
		"size":        legacyDataMatches(f, off, mid.inv, mid.paths),
		"count":       legacyDataMatches(f, more, full.inv, full.paths),
		"entry":       legacyDataMatches(f, mid.open, other, mid.paths),
		"path":        legacyDataMatches(f, mid.open, mid.inv, []string{"x"}),
		"no entry":    legacyDataMatches(f, mid.open, nil, mid.paths),
		"failing":     legacyDataMatches(&badFile{readAt: true}, mid.open, mid.inv, mid.paths),
		"zero header": legacyDataMatches(zeroReader{}, mid.open, mid.inv, mid.paths),
	} {
		if ok {
			t.Errorf("%s: want no match", name)
		}
	}
	b, _ := os.ReadFile(sealed)
	h, _ := parseObjectHeader(b[:headerSize])
	short := filepath.Join(t.TempDir(), "short")
	os.WriteFile(short, b[:h.invOff+1], 0o644)
	sf, _ := os.Open(short)
	defer sf.Close()
	b[h.invOff+h.invLen-1] ^= 0xFF // the frame checksum
	bad := filepath.Join(t.TempDir(), "bad")
	os.WriteFile(bad, b, 0o644)
	bf, _ := os.Open(bad)
	defer bf.Close()
	if legacyDataMatches(sf, mid.open, mid.inv, mid.paths) || legacyDataMatches(bf, mid.open, mid.inv, mid.paths) {
		t.Fatal("short or bad inventory: want no match")
	}
}
