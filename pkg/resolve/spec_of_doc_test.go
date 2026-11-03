//ff:func feature=resolve type=decoder control=sequence
//ff:what TestSpecOfDoc — the specs of a codec spec document's BODY; nothing from an empty document
package resolve

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestSpecOfDoc(t *testing.T) {
	doc, err := inspect.Decode(codecDocOf(0x0B80, 4, 2))
	if err != nil {
		t.Fatal(err)
	}
	if s := specOfDoc(doc); len(s) != 1 || len(s[0x0B80]) != 2 {
		t.Fatalf("%+v", s)
	}
	if s := specOfDoc(inspect.Document{}); len(s) != 0 {
		t.Fatalf("%+v", s)
	}
}
