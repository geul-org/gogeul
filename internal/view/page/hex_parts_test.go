//ff:func feature=page type=builder control=sequence
//ff:what TestHexParts — a standalone Reference names its payload fields; an ALIAS NODE splits into head lines (FORMAT noted with the name) and payload fields with the ID boxed
package page

import (
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
	"github.com/geul-org/gogeul/pkg/toulmin"
)

func TestHexParts(t *testing.T) {
	p := toulmin.NodePacket{TID: 3, Type: toulmin.NodeAlias, Span: toulmin.Span{Off: 1772, Len: 50}}
	ws, _ := toulmin.EncodeNode(p, toulmin.Wire{Span: 1})
	r, _, _ := core.DecodeRecord(ws)
	head, body := hexParts(inspect.Entry{Record: r, Words: ws, Name: "ALIAS", Data: p})
	if len(head) != 1 || head[0].Lines[0].Word != "1095" || !head[0].Lines[0].Tag || !strings.Contains(head[0].Lines[0].Note, "001001 ALIAS") {
		t.Fatalf("head: %+v", head)
	}
	if len(body) != 4 || !body[0].Record || body[0].Lines[0].Word != "0051" || body[0].Lines[1].Word != "0003" || body[2].Lines[0].Word != "06EC" {
		t.Fatalf("payload: %+v", body)
	}
	ref := []core.Word{0x02D7, 0, 0, 0xffff, 0, 0, 2, 0}
	rr, _, _ := core.DecodeRecord(ref)
	if _, body := hexParts(inspect.Entry{Record: rr, Words: ref}); len(body) != 1 || body[0].Lines[0].Name != "Issuer" || body[0].Lines[6].Name != "대상 ID" {
		t.Fatalf("reference: %+v", body)
	}
}
