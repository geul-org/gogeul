//ff:func feature=notation type=decoder control=sequence
//ff:what TestFormatCoordinate — a coordinate with a binary32 uncertainty in shortest form, a Geo one with its globe and ABSENT unit; a broken one has no short form
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestFormatCoordinate(t *testing.T) {
	second := core.Record{Kind: core.KindReference, Codec: core.CodecWikidataQID, Payload: []core.Word{0x000C, 0xA691}}
	if got, ok := formatCoordinate(core.EncodeCoordinate(core.Coordinate{Lat: -33.5, Lon: 151.25, Uncertainty: 0.1, Unit: &second})); !ok || got != "COORDINATE-EARTH(-33.5, 151.25, 0.1, Q829073)" {
		t.Fatalf("%q", got)
	}
	if got, ok := formatCoordinate(core.EncodeCoordinate(core.Coordinate{Globe: &second, Lat: 1, Lon: 2})); !ok || got != "COORDINATE-GEO(Q829073, 1, 2, 0, ABSENT)" {
		t.Fatalf("geo: %q", got)
	}
	if _, ok := formatCoordinate(core.Record{Kind: core.KindValue, Codec: core.CodecCoordinateEarth}); ok {
		t.Fatal("broken: want false")
	}
}
