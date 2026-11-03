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
	if got, err := s.setsFor(ctx, "0x00000002/documents.geul"); err != nil || !slices.Equal(got, []string{"0x00000002", ""}) {
		t.Fatal(got, err)
	}
	for _, p := range []string{issuersPath(), "h/0x00.geul", "0x00000002", "example.org/a"} {
		if got, err := s.setsFor(ctx, p); err != nil || !slices.Equal(got, []string{"", "0x00000000", "0x00000002"}) {
			t.Fatal(p, got, err)
		}
	}
}
