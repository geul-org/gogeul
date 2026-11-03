//ff:func feature=resolve type=rule control=iteration dimension=1
//ff:what TestDecodeVersionList — a list document decodes with its records and hash; odd bytes, a non-document and a bad record are ErrVerify
package resolve

import (
	"errors"
	"testing"
)

func TestDecodeVersionList(t *testing.T) {
	d := sampleKey(1, 0).DocKey()
	b := listDocOf(1, VersionEntry{Version: 1}, VersionEntry{Version: 2, Deleted: true})
	lg, err := decodeVersionList(b, d.Issuer, d.Doc)
	if err != nil || !lg.Present || len(lg.Entries) != 2 || !lg.Entries[1].Deleted || len(lg.Records) != 3 || lg.Hash != listHash(b) {
		t.Fatalf("got %+v %v", lg, err)
	}
	for i, bad := range [][]byte{{1}, {0x18, 0x0F}, listDocOf(2, VersionEntry{Version: 1})} {
		if _, err := decodeVersionList(bad, d.Issuer, d.Doc); !errors.Is(err, ErrVerify) {
			t.Fatalf("bad %d: %v", i, err)
		}
	}
}
