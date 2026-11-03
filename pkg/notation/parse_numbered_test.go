//ff:func feature=notation type=encoder control=sequence
//ff:what TestParseNumbered — #n and #label through the label map, #n it does not know as is, Qn and Pn; an unknown non-number fails
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestParseNumbered(t *testing.T) {
	ctx := Context{Local: func(l string) (uint64, bool) { return map[string]uint64{"2": 12, "edu": 20}[l], l == "2" || l == "edu" }}
	if r, _ := parseNumbered("#2", ctx); r.Payload[0] != 12 || r.Codec != core.CodecLocal {
		t.Fatalf("%+v", r)
	}
	if r, _ := parseNumbered("#edu", ctx); r.Payload[0] != 20 {
		t.Fatalf("label: %+v", r)
	}
	if r, _ := parseNumbered("#5", ctx); r.Payload[0] != 5 {
		t.Fatalf("not a label: %+v", r)
	}
	if r, _ := parseNumbered("P31", ctx); r.Payload[0] != 31 || r.Codec != core.CodecWikidataPID {
		t.Fatalf("%+v", r)
	}
	if _, err := parseNumbered("#x", ctx); err == nil {
		t.Fatal("want error")
	}
}
