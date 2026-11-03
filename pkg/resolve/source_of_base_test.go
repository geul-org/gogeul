//ff:func feature=resolve type=helper control=sequence
//ff:what TestSourceOfBase — file:// and https:// bases open a Source; an empty base is ErrNotFound and a bad entry fails
package resolve

import (
	"errors"
	"testing"
)

func TestSourceOfBase(t *testing.T) {
	if s, err := sourceOfBase("file://" + t.TempDir()); err != nil || s == nil {
		t.Fatal(s, err)
	}
	if _, err := sourceOfBase(""); !errors.Is(err, ErrNotFound) {
		t.Fatal("empty: want ErrNotFound", err)
	}
	if _, err := sourceOfBase("http://geul.org"); err == nil {
		t.Fatal("http: want error")
	}
}
