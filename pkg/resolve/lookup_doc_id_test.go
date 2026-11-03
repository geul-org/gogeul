//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what TestLookupDocID — a 16-digit document ID and an 8-digit codec ID are read; a document name and other spellings are ErrNotFound naming the Issuer
package resolve

import (
	"errors"
	"slices"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestLookupDocID(t *testing.T) {
	g := []core.Word{0, 0}
	if id, err := lookupDocID(g, "0xFFFF000000000007"); err != nil || !slices.Equal(id, []core.Word{0xffff, 0, 0, 7}) {
		t.Fatal(id, err)
	}
	if id, err := lookupDocID(g, "0x00000B80"); err != nil || !slices.Equal(id, []core.Word{0, 0x0B80}) {
		t.Fatal(id, err)
	}
	if id, err := lookupDocID(g, "0xffff000000000007"); err != nil || !slices.Equal(id, []core.Word{0xffff, 0, 0, 7}) {
		t.Fatal("lower case:", id, err)
	}
	for _, s := range []string{"seven", "0xB80", "0XFFFF000000000007", ""} {
		if _, err := lookupDocID(g, s); !errors.Is(err, ErrNotFound) || !strings.Contains(err.Error(), "0x00000000") {
			t.Fatal(s, err)
		}
	}
}
