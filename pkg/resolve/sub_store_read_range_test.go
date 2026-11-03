//ff:func feature=resolve type=resolver control=sequence
//ff:what TestSubStoreReadRange — a range of the set's own object
package resolve

import (
	"context"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
)

func TestSubStoreReadRange(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	d.write("0x00000002/o/a.gpk", []byte("abcdef"))
	s := subStore{base: pack.FileStore{Root: d.file("")}, name: "0x00000002"}
	if b, err := s.ReadRange(context.Background(), "o/a.gpk", 2, 3); err != nil || string(b) != "cde" {
		t.Fatal(string(b), err)
	}
}
