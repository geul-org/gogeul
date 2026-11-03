//ff:func feature=build type=helper control=sequence
//ff:what TestTreeName — URL paths under /.well-known/geul map to fs names, the top to "."
package build

import "testing"

func TestTreeName(t *testing.T) {
	if treeName("/.well-known/geul") != "." || treeName("/.well-known/geul/") != "." || treeName("/.well-known/geul/h/a.geul") != "h/a.geul" {
		t.Fatal("treeName")
	}
}
