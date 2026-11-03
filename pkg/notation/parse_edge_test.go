//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what TestParseEdge — label, name and fields by spec; no label; a whole raw Edge; a bad label, no name, an unclosed quote fail
package notation

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestParseEdge(t *testing.T) {
	ctx := Context{Spec: tripleSpec(false)}
	e, err := ParseEdge("#4 TRIPLE P31 Q42 Q5 preferred", ctx)
	if err != nil || e.Label != "4" || e.Codec != 0x0A || !slices.Equal(e.Rest, []core.Word{0x0191, 31, 0x0141, 42, 0x0141, 5, 1}) {
		t.Fatalf("%+v %v", e, err)
	}
	if e, err := ParseEdge(`label Q42 "x"`, ctx); err != nil || e.Label != "" || e.Codec != 0x04 {
		t.Fatalf("%+v %v", e, err)
	}
	if e, err := ParseEdge("0x0a43005100070001", ctx); err != nil || e.Label != "7" || e.Codec != 0x04 {
		t.Fatalf("raw: %+v %v", e, err)
	}
	if e, err := ParseEdge("#edu TRIPLE P31 Q42 Q5 preferred", ctx); err != nil || e.Label != "edu" {
		t.Fatalf("named label: %+v %v", e, err)
	}
	for _, s := range []string{"#x TRIPLE", "#a.b TRIPLE P31 Q42 Q5 normal", "# TRIPLE P31 Q42 Q5 normal", "#1", `LABEL "open`, "NOPE Q1", "TRIPLE P31", "LABEL word", "0x00510001"} {
		if _, err := ParseEdge(s, ctx); err == nil {
			t.Errorf("%s: want error", s)
		}
	}
}
