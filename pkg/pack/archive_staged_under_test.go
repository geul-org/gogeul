//ff:func feature=pack type=resolver control=sequence
//ff:what TestArchiveStagedUnder — a staged write under the prefix; nothing staged there
package pack

import "testing"

func TestArchiveStagedUnder(t *testing.T) {
	a := &Archive{puts: map[string][]byte{"a/b.geul": {}}}
	if !a.stagedUnder("a/") || a.stagedUnder("b/") {
		t.Fatal("stagedUnder")
	}
}
