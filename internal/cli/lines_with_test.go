//ff:func feature=cli type=helper control=sequence
//ff:what TestLinesWith — a record laid out by the given spec; a document that does not decode fails
package cli

import (
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
)

func TestLinesWith(t *testing.T) {
	label := core.Record{Kind: core.KindEdge, Codec: 0x04, Payload: []core.Word{0x0088, 1, 0x00A0, 42}}
	b, _ := core.EncodeDocument([]core.Record{core.EncodeSpecVersion(2)}, []core.Record{label})
	ls, err := linesWith(b, codecspec.Spec{})
	if err != nil || len(ls.Head) != 1 || len(ls.Body) != 1 || !strings.HasPrefix(ls.Body[0], "#1 LABEL") {
		t.Fatalf("%+v %v", ls, err)
	}
	if _, err := linesWith([]byte{1}, nil); err == nil {
		t.Fatal("want error")
	}
}
