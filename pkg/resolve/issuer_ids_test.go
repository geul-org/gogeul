//ff:func feature=resolve type=decoder control=iteration dimension=1
//ff:what TestIssuerIDs — the bound IDs; nothing for an absent or unreadable index
package resolve

import (
	"context"
	"slices"
	"testing"
)

func TestIssuerIDs(t *testing.T) {
	b, err := formsTree(t).Fetch(context.Background(), issuersPath())
	if err != nil {
		t.Fatal(err)
	}
	got := issuerIDs(b)
	slices.Sort(got)
	if !slices.Equal(got, []string{"0x00000000", "0x00000002"}) {
		t.Fatal(got)
	}
	for _, junk := range [][]byte{nil, []byte("x")} {
		if got := issuerIDs(junk); got != nil {
			t.Fatal(got)
		}
	}
}
