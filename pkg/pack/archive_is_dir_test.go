//ff:func feature=pack type=resolver control=sequence
//ff:what TestArchiveIsDir — the top, folders holding paths; files and absent names are not
package pack

import "testing"

func TestArchiveIsDir(t *testing.T) {
	a, _ := testArchive(t)
	if !a.isDir(".") || !a.isDir("0x00000000") || !a.isDir("0x00000000/hashes") || !a.isDir("issuers") || a.isDir("h") || a.isDir("issuers/0x00000001.geul") || a.isDir("zz") {
		t.Fatal("isDir")
	}
}
