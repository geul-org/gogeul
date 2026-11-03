//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what TestIssuerOfID — 8, 16 and 64 digits read as one Issuer at its issued width; names and other widths are ErrNotFound
package resolve

import (
	"errors"
	"slices"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestIssuerOfID(t *testing.T) {
	for _, s := range []string{"0x00000005", "0x0000000000000005", "0x" + strings.Repeat("0", 63) + "5"} {
		if ws, err := issuerOfID(s); err != nil || !slices.Equal(ws, []core.Word{0, 5}) {
			t.Fatal(s, ws, err)
		}
	}
	for _, s := range []string{"geul.org", "0x0005", "0x0000000G"} {
		if _, err := issuerOfID(s); !errors.Is(err, ErrNotFound) {
			t.Fatal(s, err)
		}
	}
}
