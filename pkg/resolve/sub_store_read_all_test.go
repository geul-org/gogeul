//ff:func feature=resolve type=resolver control=sequence
//ff:what TestSubStoreReadAll — reads the set's own object, not the place's
package resolve

import (
	"context"
	"errors"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
)

func TestSubStoreReadAll(t *testing.T) {
	root := t.TempDir()
	d := Dir{Root: root}
	d.write("latest.txt", []byte("top"))
	d.write("0x0001/latest.txt", []byte("one"))
	base := pack.FileStore{Root: d.file("")}
	ctx := context.Background()
	if b, err := (subStore{base: base, name: "0x0001"}).ReadAll(ctx, "latest.txt"); err != nil || string(b) != "one" {
		t.Fatal(string(b), err)
	}
	if _, err := (subStore{base: base, name: "0x0002"}).ReadAll(ctx, "latest.txt"); !errors.Is(err, ErrNotFound) {
		t.Fatal(err)
	}
}
