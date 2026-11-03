//ff:func feature=resolve type=helper control=sequence
//ff:what TestDocIDOf — the DOC-ID of a named document; false for an unnamed one
package resolve

import (
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
	"testing"
)

func TestDocIDOf(t *testing.T) {
	doc, _ := inspect.Decode(sampleDoc(2, "x"))
	if n, ok := docIDOf(doc); !ok || core.FormatID(n.Doc) != "0xFFFF000000000002" {
		t.Fatalf("named: %+v %v", n, ok)
	}
	if _, ok := docIDOf(inspect.Document{}); ok {
		t.Fatal("unnamed: want false")
	}
}
