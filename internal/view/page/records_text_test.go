//ff:func feature=page type=helper control=sequence
//ff:what TestRecordsText — several records joined by spaces; words that do not decode end as one hex run
package page

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestRecordsText(t *testing.T) {
	a, _ := core.EncodeRecord(core.EncodeBytesValue(core.CodecURIEven, []byte("https://x")))
	b, _ := core.EncodeRecord(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("v1")))
	if got := recordsText(slices.Concat(a, b)); got != `https://x "v1"` {
		t.Fatalf("got %q", got)
	}
	if got := recordsText([]core.Word{0xFFFF}); got != "0xFFFF" {
		t.Fatalf("bad: %q", got)
	}
}
