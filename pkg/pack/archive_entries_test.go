//ff:func feature=pack type=resolver control=sequence
//ff:what TestArchiveEntries — entries carry their full path and Archive; a file's Info has its real size
package pack

import "testing"

func TestArchiveEntries(t *testing.T) {
	a, _ := testArchive(t)
	idx, _ := a.index()
	if top := a.entries(idx, "."); len(top) != 2 || top[1].path != "issuers" || !top[1].IsDir() {
		t.Fatal(top)
	}
	es := a.entries(idx, "issuers")
	if len(es) != 1 || es[0].path != "issuers/0x00000001.geul" || es[0].a != a {
		t.Fatal(es)
	}
	fi, err := es[0].Info()
	if err != nil || fi.Size() != int64(len(fixtureFile(t, "issuers/0x00000001.geul"))) {
		t.Fatal(err)
	}
}
