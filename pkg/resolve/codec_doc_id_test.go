//ff:func feature=resolve type=helper control=sequence
//ff:what TestCodecDocID — the target word widened to 4 words
package resolve

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCodecDocID(t *testing.T) {
	if got := codecDocID(0x03D0); !slices.Equal(got, []core.Word{0, 0x03D0}) {
		t.Fatalf("%04X", got)
	}
}
