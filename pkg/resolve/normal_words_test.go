//ff:func feature=resolve type=helper control=sequence
//ff:what TestNormalWords — a padded Issuer folds to its issued width; a 1-word Issuer stays as it is
package resolve

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestNormalWords(t *testing.T) {
	if got := normalWords([]core.Word{0, 0, 0, 5}); !slices.Equal(got, []core.Word{0, 5}) {
		t.Fatal(got)
	}
	if got := normalWords([]core.Word{5}); !slices.Equal(got, []core.Word{5}) {
		t.Fatal(got)
	}
}
