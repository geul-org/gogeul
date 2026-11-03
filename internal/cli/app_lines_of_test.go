//ff:func feature=cli type=helper control=sequence
//ff:what TestAppLinesOf — HEAD and BODY lines; a document that does not decode fails
package cli

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestAppLinesOf(t *testing.T) {
	a := &app{}
	label := core.Record{Kind: core.KindEdge, Codec: 0x04, Payload: []core.Word{0x0051, 1, 0x0141, 42}}
	b, _ := core.EncodeDocument([]core.Record{core.EncodeSpecVersion(2)}, []core.Record{label})
	ls, err := a.linesOf(b)
	if err != nil || len(ls.Head) != 1 || len(ls.Body) != 1 || ls.Body[0] != "#1 LABEL Q42" {
		t.Fatalf("%+v %v", ls, err)
	}
	if _, err := a.linesOf([]byte{1}); err == nil {
		t.Fatal("want error")
	}
}
