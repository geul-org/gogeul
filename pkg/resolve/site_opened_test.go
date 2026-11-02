//ff:func feature=resolve type=resolver control=sequence
//ff:what TestSiteOpened — only sets that exist are listed
package resolve

import (
	"context"
	"slices"
	"testing"
)

func TestSiteOpened(t *testing.T) {
	s := localSite(packedRoot(t, formsTree(t), false, nil))
	ctx := context.Background()
	if got := s.opened(); len(got) != 0 {
		t.Fatal(got)
	}
	s.set(ctx, "0x0000")
	s.set(ctx, "")
	if got := s.opened(); !slices.Equal(got, []string{"0x0000"}) {
		t.Fatal(got)
	}
}
