//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what TestRecordField — an allowed codec, any record when none are declared; slot markers only where named exactly; a disallowed codec and a bad token fail
package notation

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
)

func TestRecordField(t *testing.T) {
	qid := codecspec.Field{Codecs: []core.Word{core.TargetWord(core.KindReference, core.CodecWikidataQID)}}
	if _, err := recordField("Q5", qid, Context{}); err != nil {
		t.Fatal(err)
	}
	if _, err := recordField("P5", qid, Context{}); err == nil {
		t.Fatal("disallowed: want error")
	}
	if _, err := recordField("word", codecspec.Field{}, Context{}); err == nil {
		t.Fatal("bad token: want error")
	}
	if ws, err := recordField("P5", codecspec.Field{}, Context{}); err != nil || len(ws) != 2 {
		t.Fatalf("no CODEC, any record: %v %v", ws, err)
	}
	object := codecspec.Field{Codecs: []core.Word{0x0780, 0x0F80, core.UnknownValue, core.NoValue}}
	if ws, err := recordField("UNKNOWN", object, Context{}); err != nil || !slices.Equal(ws, []core.Word{core.UnknownValue}) {
		t.Fatalf("UNKNOWN named: %v %v", ws, err)
	}
	if ws, err := recordField("NONE", object, Context{}); err != nil || !slices.Equal(ws, []core.Word{core.NoValue}) {
		t.Fatalf("NONE named: %v %v", ws, err)
	}
	for _, tok := range []string{"ABSENT", "0x0003"} {
		if _, err := recordField(tok, object, Context{}); err == nil {
			t.Fatalf("%s not named, wildcards do not take it: want error", tok)
		}
	}
	if _, err := recordField("UNKNOWN", codecspec.Field{}, Context{}); err == nil {
		t.Fatal("marker in a field without CODECs: want error")
	}
}
