//ff:func feature=cli type=helper control=sequence
//ff:what TestAddressesOf — Issuer/document for each
package cli

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/resolve"
)

func TestAddressesOf(t *testing.T) {
	if got := addressesOf("0x00000000", []resolve.Named{{ID: "0x01"}, {ID: "0x02"}}); !slices.Equal(got, []string{"0x00000000/0x01", "0x00000000/0x02"}) {
		t.Fatal(got)
	}
}
