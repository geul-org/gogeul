//ff:func feature=notation type=decoder control=iteration dimension=1
//ff:what TestFormatEntry — Edge and Meta lines by layout; unknown names, Issuer records and Edges without a LOCAL Edge Id are raw words
package notation

import (
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestFormatEntry(t *testing.T) {
	doc, _ := inspect.Decode(sampleDoc())
	if got := FormatEntry(doc.Head[0]); got != "SPEC-VERSION 0x0002" {
		t.Fatal(got)
	}
	if got := FormatEntry(doc.Body[0]); got != `#1 LABEL Q42 "Douglas Adams"@en` {
		t.Fatal(got)
	}
	if got := FormatEntry(doc.Body[1]); got != "#2 TRIPLE 0x0191001F0141002A014100050001" {
		t.Fatal(got)
	}
	unnamed := doc.Body[0]
	unnamed.Name = ""
	noID := inspect.Entry{Record: core.Record{Kind: core.KindEdge, Codec: 0x04, Payload: []core.Word{0x0141, 42}}, Words: []core.Word{0x0A42, 0x0141, 42}, Name: "LABEL"}
	issued := doc.Head[0]
	issued.Record.Format = 1
	for _, e := range []inspect.Entry{unnamed, noID, issued} {
		if got := FormatEntry(e); !strings.HasPrefix(got, "0x") {
			t.Errorf("want raw, got %s", got)
		}
	}
}
