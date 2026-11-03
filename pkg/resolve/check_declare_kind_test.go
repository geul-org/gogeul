//ff:func feature=resolve type=rule control=sequence
//ff:what TestCheckDeclareKind — a codec document of the record's Kind passes; another Kind fails (§1.1 103); a document without DECLARE passes
package resolve

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestCheckDeclareKind(t *testing.T) {
	edge, _ := inspect.Decode(codecDocOf(0x0980, 4, 2))
	if err := checkDeclareKind(edge, core.KindEdge); err != nil {
		t.Fatal(err)
	}
	if err := checkDeclareKind(edge, core.KindValue); err == nil {
		t.Fatal("Edge codec for a Value record: want error")
	}
	value, _ := inspect.Decode(codecDocOf(0x05A0, 4))
	if err := checkDeclareKind(value, core.KindEdge); err == nil {
		t.Fatal("Value codec for an Edge record: want error")
	}
	none, _ := inspect.Decode(edgeDoc(7, "a"))
	if err := checkDeclareKind(none, core.KindValue); err != nil {
		t.Fatal(err)
	}
}
