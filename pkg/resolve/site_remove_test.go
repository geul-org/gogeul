//ff:func feature=resolve type=resolver control=sequence
//ff:what TestSiteRemove — a source refuses removals
package resolve

import "testing"

func TestSiteRemove(t *testing.T) {
	if err := localSite(t.TempDir()).Remove("a"); err == nil {
		t.Fatal("remove")
	}
}
