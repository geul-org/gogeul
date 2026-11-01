//ff:func feature=pack type=resolver control=sequence
//ff:what TestArchiveIsDirPaths — isDir agrees with the path list on every folder, file and absent name, across staged writes, staged removals and committed deletion markers
package pack

import (
	"context"
	"testing"
)

func TestArchiveIsDirPaths(t *testing.T) {
	a, _ := testArchive(t)
	checkIsDir(t, a, "opened")
	doc := "0x0000/0x0000000000000008"
	a.WriteFile("stage/sub/new.geul", []byte("n"))
	a.Remove(doc + "/0x00000001.geul")
	checkIsDir(t, a, "staged")
	a.Remove(doc + "/versions.geul")
	checkIsDir(t, a, "doc removed, staged")
	if _, err := a.Commit(context.Background()); err != nil {
		t.Fatal(err)
	}
	a.idx = nil
	checkIsDir(t, a, "committed")
	if a.isDir(doc) {
		t.Fatal("a removed document is not a directory")
	}
}
