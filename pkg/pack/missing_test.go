//ff:func feature=pack type=helper control=sequence
//ff:what TestMissing — fs.ErrNotExist becomes ErrNotFound, other errors pass through
package pack

import (
	"errors"
	"io/fs"
	"testing"
)

func TestMissing(t *testing.T) {
	if !errors.Is(missing(fs.ErrNotExist, "x"), ErrNotFound) {
		t.Fatal("not exist")
	}
	other := errors.New("other")
	if missing(other, "x") != other {
		t.Fatal("other")
	}
}
