//ff:func feature=pack type=resolver control=sequence
//ff:what TestReaderReadRange — bytes of a listed object; a missing object is ErrCorrupt, not NotFound
package pack

import (
	"context"
	"errors"
	"testing"
)

func TestReaderReadRange(t *testing.T) {
	obj, o := objBytes(kindData, []byte("ab"), nil)
	r := newReader(memStore{objectName(o): obj}, Options{}, &manifest{}, "")
	if b, err := r.readRange(context.Background(), o, 64, 2); err != nil || string(b) != "ab" {
		t.Fatal(err)
	}
	_, err := r.readRange(context.Background(), object{kind: kindData}, 64, 2)
	if !errors.Is(err, ErrCorrupt) || errors.Is(err, ErrNotFound) {
		t.Fatal(err)
	}
}
