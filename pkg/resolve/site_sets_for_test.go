//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestSiteSetsFor — `<Issuer>/…` asks its set and the top set; h/, issuers.geul and a bare name ask every set
package resolve

import (
	"context"
	"slices"
	"testing"
)

func TestSiteSetsFor(t *testing.T) {
	s := localSite(packedRoot(t, formsTree(t), false, nil))
	ctx := context.Background()
	if got, err := s.setsFor(ctx, "0x0001/documents.geul"); err != nil || !slices.Equal(got, []string{"0x0001", ""}) {
		t.Fatal(got, err)
	}
	for _, p := range []string{issuersFile, "h/0x00.geul", "0x0001", "example.org/a"} {
		if got, err := s.setsFor(ctx, p); err != nil || !slices.Equal(got, []string{"", "0x0000", "0x0001"}) {
			t.Fatal(p, got, err)
		}
	}
}
