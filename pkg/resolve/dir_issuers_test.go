//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirIssuers — Issuer directories by width then ID with their names; other entries skipped; no archive fails
package resolve

import (
	"os"
	"testing"
)

func TestDirIssuers(t *testing.T) {
	d := issuersFixture(t)
	os.MkdirAll(d.file("0x00000001"), 0o755)
	os.MkdirAll(d.file("other"), 0o755)
	got, err := d.Issuers()
	if err != nil || len(got) != 2 || got[0].ID != "0x00000000" || got[0].Names[0] != "geul.org" || got[1].ID != "0x00000001" || len(got[1].Names) != 0 {
		t.Fatalf("%+v %v", got, err)
	}
	if _, err := (Dir{Root: t.TempDir()}).Issuers(); err == nil {
		t.Fatal("no archive: want error")
	}
}
