//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirAt — a version by number, 0 for the newest live; an unknown or deleted version and a broken list fail
package resolve

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDirAt(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	issuer, doc := []core.Word{0, 0}, []core.Word{0xffff, 0, 0, 7}
	v1, v2 := edgeDoc(7, "a"), edgeDoc(7, "a", "b")
	d.Publish(v1)
	d.Publish(v2)
	if b, err := d.At(issuer, doc, 1); err != nil || !bytes.Equal(b, v1) {
		t.Fatalf("v1: %v", err)
	}
	if b, err := d.At(issuer, doc, 0); err != nil || !bytes.Equal(b, v2) {
		t.Fatalf("newest: %v", err)
	}
	d.Tombstone(issuer, doc, 1, 1)
	if _, err := d.At(issuer, doc, 1); !errors.Is(err, ErrNotFound) || !strings.Contains(err.Error(), "deleted") {
		t.Fatalf("deleted: %v", err)
	}
	if _, err := d.At(issuer, doc, 9); !errors.Is(err, ErrNotFound) || strings.Contains(err.Error(), "deleted") {
		t.Fatalf("unknown: %v", err)
	}
	d.write(listPath(sampleKeyOf(7)), []byte{1})
	if _, err := d.At(issuer, doc, 0); err == nil {
		t.Fatal("broken list: want error")
	}
}
