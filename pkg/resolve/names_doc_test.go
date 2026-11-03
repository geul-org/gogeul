//ff:func feature=resolve type=encoder control=sequence
//ff:what TestNamesDoc — a new index gets SPEC-VERSION and Edge Id 1; the next BIND follows the highest; a bad Subject fails
package resolve

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestNamesDoc(t *testing.T) {
	subj, _ := core.EncodeDocPointer([]core.Word{0, 0}, []core.Word{0, 0, 0, 7}, []core.Word{0, 0})
	b, err := namesDoc(nil, nil, core.Bind{EdgeID: 1, Subject: subj, Name: []byte("a")})
	if err != nil {
		t.Fatal(err)
	}
	ws, _ := core.ReadWords(b)
	doc, _ := core.SplitDocument(ws)
	b2, _ := namesDoc(doc.Head, doc.Body, core.Bind{EdgeID: 1, Subject: subj, Name: []byte("b")})
	ws, _ = core.ReadWords(b2)
	doc2, _ := core.SplitDocument(ws)
	last, _ := core.DecodeBind(doc2.Body[1])
	if len(doc.Head) != 1 || len(doc2.Body) != 2 || last.EdgeID != 2 {
		t.Fatalf("head %d body %d id %d", len(doc.Head), len(doc2.Body), last.EdgeID)
	}
	if _, err := namesDoc(nil, nil, core.Bind{EdgeID: 1, Subject: core.Record{Kind: core.KindEdge, Codec: 0x140}}); err == nil {
		t.Fatal("bad subject: want error")
	}
}
