//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what checkIsDir — test fixture: isDir against the path list on every path, its folders and some absent names
package pack

import (
	"path"
	"testing"
)

// checkIsDir fails when isDir and a.Paths disagree; stage names the moment.
func checkIsDir(t *testing.T, a *Archive, stage string) {
	t.Helper()
	ps, err := a.Paths()
	if err != nil {
		t.Fatal(err)
	}
	names := []string{"zz", "0x00000002", "0x00000002/docs", "h/zz", "0x00000000/zz", "0x00000000/docs/0x00000000000000FF", "0x00000000/docs/0x0000000000000008/zz", "0x00000000/hashes", "issuers", "stage"}
	for _, p := range ps {
		names = append(names, p, path.Dir(p), path.Dir(path.Dir(p)))
	}
	for _, n := range names {
		if want := n == "." || hasPrefix(ps, dirPrefix(n)); a.isDir(n) != want {
			t.Fatalf("%s: isDir(%s) = %v, listing says %v", stage, n, !want, want)
		}
	}
}
