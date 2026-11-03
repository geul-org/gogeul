//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what TestParseCoordinate — Earth and Geo read back from their short forms bit for bit; ABSENT unit; wrong arity, an unreadable number, globe or unit, and a unit that is not a Reference fail
package notation

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestParseCoordinate(t *testing.T) {
	r, err := parseCall("COORDINATE-EARTH(52.516666666667, 13.383333333333, 1, Q209426)", Context{})
	ws, _ := core.EncodeRecord(r)
	if err != nil || !slices.Equal(ws, []core.Word{0x044D, 0x404A, 0x4222, 0x2222, 0x2251, 0x402A, 0xC444, 0x4444, 0x4389, 0x3F80, 0x0000, 0x0142, 0x0003, 0x3212}) {
		t.Fatalf("earth: %04X %v", ws, err)
	}
	if got, ok := formatValue(r); !ok || got != "COORDINATE-EARTH(52.516666666667, 13.383333333333, 1, Q209426)" {
		t.Fatalf("earth back: %q", got)
	}
	r, err = parseCall("COORDINATE-GEO(Q111, 4.5, 137.4, 0, ABSENT)", Context{})
	if c, _ := core.DecodeCoordinate(r); err != nil || c.Globe == nil || c.Unit != nil || c.Lon != 137.4 {
		t.Fatalf("geo: %+v %v", c, err)
	}
	if got, ok := formatValue(r); !ok || got != "COORDINATE-GEO(Q111, 4.5, 137.4, 0, ABSENT)" {
		t.Fatalf("geo back: %q", got)
	}
	for _, tok := range []string{"COORDINATE-EARTH(1, 2, 3)", "COORDINATE-EARTH(x, 2, 0, Q1)", "COORDINATE-EARTH(1, 2, 0, word)",
		"COORDINATE-GEO(word, 1, 2, 0, Q1)", `COORDINATE-EARTH(1, 2, 0, "m")`} {
		if _, err := parseCall(tok, Context{}); err == nil {
			t.Errorf("%s: want error", tok)
		}
	}
}
