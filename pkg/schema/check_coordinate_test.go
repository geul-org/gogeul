//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestCheckCoordinate — valid Earth and Geo points pass (with a unit, exact, or unknown); bad width, no globe, a quantified, WIKIDATA-PID or ISSUER-POINTER globe, non-finite and out-of-range degrees, a negative, -0, infinite or NaN uncertainty, a quantified unit and a non-zero uncertainty without a unit fail
package schema

import (
	"math"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCheckCoordinate(t *testing.T) {
	q := func(n core.Word, quant uint16) core.Record {
		return core.Record{Kind: core.KindReference, Codec: core.CodecWikidataQID + quant, Payload: []core.Word{n}}
	}
	metre, someUnit := q(11573, 0), q(11573, 1)
	geo := func(g core.Record, lat, lon float64) core.Record {
		return core.EncodeCoordinate(core.Coordinate{Globe: &g, Lat: lat, Lon: lon})
	}
	cases := []struct {
		r    core.Record
		want string
	}{
		{core.EncodeCoordinate(core.Coordinate{Lat: 52.5, Lon: 13.4, Uncertainty: 4.4, Unit: &metre}), ""},
		{core.EncodeCoordinate(core.Coordinate{Lat: 52.5, Lon: 13.4}), ""},
		{core.EncodeCoordinate(core.Coordinate{Lat: 52.5, Lon: 13.4, Unit: &metre}), ""},
		{core.EncodeCoordinate(core.Coordinate{Lat: 52.5, Lon: 13.4, Uncertainty: 1}), "V8"},
		{core.EncodeCoordinate(core.Coordinate{Lat: 52.5, Lon: 13.4, Uncertainty: -1, Unit: &metre}), "V8"},
		{core.EncodeCoordinate(core.Coordinate{Lat: 52.5, Lon: 13.4, Uncertainty: float32(math.Copysign(0, -1)), Unit: &metre}), "V8"},
		{core.EncodeCoordinate(core.Coordinate{Lat: 52.5, Lon: 13.4, Uncertainty: float32(math.Inf(1)), Unit: &metre}), "V8"},
		{core.EncodeCoordinate(core.Coordinate{Lat: 52.5, Lon: 13.4, Uncertainty: float32(math.NaN()), Unit: &metre}), "V8"},
		{core.EncodeCoordinate(core.Coordinate{Lat: 52.5, Lon: 13.4, Uncertainty: 1, Unit: &someUnit}), "V8"},
		{geo(q(111, 0), 4.5, 359), ""},
		{core.Record{Kind: core.KindValue, Codec: 0x04, Payload: make([]core.Word, 10)}, "V8"},
		{geo(q(111, 1), 4.5, 10), "V8"},
		{geo(core.Record{Kind: core.KindReference, Codec: core.CodecWikidataQID + 5, Payload: []core.Word{5}}, 4.5, 10), "V8"},
		{geo(core.Record{Kind: core.KindReference, Codec: 0x40, Payload: []core.Word{0, 1}}, 4.5, 10), "V8"},
		{geo(core.Record{Kind: core.KindReference, Codec: 0x05, Payload: []core.Word{7}}, 4.5, 10), ""},
		{core.EncodeCoordinate(core.Coordinate{Lat: math.NaN(), Lon: 1}), "V8"},
		{core.EncodeCoordinate(core.Coordinate{Lat: 1, Lon: math.Inf(1)}), "V8"},
		{core.EncodeCoordinate(core.Coordinate{Lat: 91, Lon: 1}), "V8"},
		{core.EncodeCoordinate(core.Coordinate{Lat: 1, Lon: 200}), "V8"},
		{geo(q(111, 0), 1, 361), "V8"},
		{geo(q(111, 0), -10, -340), ""},
		{geo(q(111, 0), 1, -360), ""},
		{geo(q(111, 0), 1, 360), ""},
		{geo(q(111, 0), 1, -360.000001), "V8"},
		{geo(q(111, 0), 1, 360.000001), "V8"},
		{core.EncodeCoordinate(core.Coordinate{Lat: 1, Lon: -180}), ""},
		{core.EncodeCoordinate(core.Coordinate{Lat: 1, Lon: 180}), ""},
		{core.EncodeCoordinate(core.Coordinate{Lat: 1, Lon: -181}), "V8"},
	}
	for i, c := range cases {
		if got := rulesOf(checkCoordinate(recOf(c.r))); got != c.want {
			t.Fatalf("case %d: %q want %q", i, got, c.want)
		}
	}
}
