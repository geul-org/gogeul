//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what TestParseFields — tokens per field, a rest field taken or left out; missing, extra and bad tokens fail
package notation

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestParseFields(t *testing.T) {
	fields := tripleSpec(false)[core.TargetWord(core.KindEdge, 0x0A)]
	ws, err := parseFields([]string{"P31", "Q42", "Q5", "preferred"}, fields, Context{})
	if err != nil || !slices.Equal(ws, []core.Word{0x0191, 31, 0x0141, 42, 0x0141, 5, 1}) {
		t.Fatalf("%04X %v", ws, err)
	}
	withRest := tripleSpec(true)[core.TargetWord(core.KindEdge, 0x0A)]
	if ws, err := parseFields([]string{"P31", "Q42", "Q5", "normal"}, withRest, Context{}); err != nil || len(ws) != 7 {
		t.Fatalf("rest left out: %v", err)
	}
	if ws, err := parseFields([]string{"P31", "Q42", "Q5", "normal", "0x00010002"}, withRest, Context{}); err != nil || len(ws) != 9 {
		t.Fatalf("rest given: %v", err)
	}
	for _, toks := range [][]string{{"P31", "Q42"}, {"P31", "Q42", "Q5", "normal", "extra"}, {"Q1", "Q42", "Q5", "normal"}} {
		if _, err := parseFields(toks, fields, Context{}); err == nil {
			t.Errorf("%v: want error", toks)
		}
	}
}
