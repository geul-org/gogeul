//ff:func feature=resolve type=resolver control=sequence
//ff:what TestSiteWriteFile — a source refuses writes and a Dir over it cannot publish
package resolve

import "testing"

func TestSiteWriteFile(t *testing.T) {
	s := localSite(t.TempDir())
	if err := s.WriteFile("a", nil); err == nil {
		t.Fatal("write")
	}
	if _, err := (Dir{Store: s}).Publish(edgeDoc(7, "a")); err == nil {
		t.Fatal("publish")
	}
}
