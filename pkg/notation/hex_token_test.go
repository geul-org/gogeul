//ff:func feature=notation type=decoder control=sequence
//ff:what TestHexToken — four upper-case hex digits per word after 0x
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestHexToken(t *testing.T) {
	if got := hexToken([]core.Word{0x0B86, 0x00ff}); got != "0x0B8600FF" {
		t.Fatal(got)
	}
}
