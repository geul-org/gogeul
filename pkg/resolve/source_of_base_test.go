//ff:func feature=resolve type=helper control=sequence
//ff:what TestSourceOfBase — file:// and https:// bases open a Source; an empty base and a bad entry fail
package resolve

import "testing"

func TestSourceOfBase(t *testing.T) {
	if s, err := sourceOfBase("file://" + t.TempDir()); err != nil || s == nil {
		t.Fatal(s, err)
	}
	if _, err := sourceOfBase(""); err == nil {
		t.Fatal("empty: want error")
	}
	if _, err := sourceOfBase("http://geul.org"); err == nil {
		t.Fatal("http: want error")
	}
}
