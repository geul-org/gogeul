//ff:func feature=notation type=encoder control=sequence
//ff:what TestParseFraction — FRACTION(2, 3) gives the grammar's words and is not reduced; wrong arity and a bad component fail
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestParseFraction(t *testing.T) {
	r, err := parseFraction([]string{"2", "4"})
	ws, _ := core.EncodeRecord(r)
	if err != nil || core.FormatID(ws) != "0x04740411000204210004" {
		t.Fatalf("2/4: %s %v", core.FormatID(ws), err)
	}
	if _, err := parseFraction([]string{"2"}); err == nil {
		t.Fatal("arity: want error")
	}
	if _, err := parseFraction([]string{"2", "-3"}); err == nil {
		t.Fatal("negative denominator: want error")
	}
}
