//ff:func feature=notation type=decoder control=iteration dimension=1
//ff:what TestMarkerToken — the three named slot markers have tokens; a reserved marker and other records have none
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestMarkerToken(t *testing.T) {
	for w, want := range map[core.Word]string{core.Absent: "ABSENT", core.UnknownValue: "UNKNOWN", core.NoValue: "NONE"} {
		r, _, _ := core.DecodeRecord([]core.Word{w})
		if got, ok := markerToken(r); !ok || got != want {
			t.Fatalf("0x%04X: %q %v", w, got, ok)
		}
	}
	reserved, _, _ := core.DecodeRecord([]core.Word{0x0011, 7})
	if _, ok := markerToken(reserved); ok {
		t.Fatal("reserved marker: want false")
	}
	if _, ok := markerToken(core.EncodeLocalRef(1)); ok {
		t.Fatal("LOCAL: want false")
	}
}
