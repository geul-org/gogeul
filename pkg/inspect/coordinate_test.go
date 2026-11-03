//ff:func feature=inspect type=decoder control=sequence
//ff:what TestCoordinate — Earth has no body and its unit as Qn, an unknown uncertainty no unit, a QID body reads as Qn, another Reference by codec and hex; a broken value fails
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCoordinate(t *testing.T) {
	degree := core.Record{Kind: core.KindReference, Codec: core.CodecWikidataQID, Payload: []core.Word{28390}}
	v, err := coordinate(core.EncodeCoordinate(core.Coordinate{Lat: 52.5, Lon: 13.4, Uncertainty: 0.25, Unit: &degree}))
	if err != nil || v.Codec != core.CodecCoordinateEarth || v.Lat != 52.5 || v.Lon != 13.4 || v.Float != 0.25 || v.Unit != "Q28390" || v.Globe != "" {
		t.Fatalf("earth: %+v %v", v, err)
	}
	mars := core.Record{Kind: core.KindReference, Codec: core.CodecWikidataQID, Payload: []core.Word{111}}
	if v, err := coordinate(core.EncodeCoordinate(core.Coordinate{Globe: &mars, Lat: 1, Lon: 2})); err != nil || v.Globe != "Q111" || v.Unit != "" {
		t.Fatalf("mars: %+v %v", v, err)
	}
	local := core.EncodeLocalRef(3)
	if v, _ := coordinate(core.EncodeCoordinate(core.Coordinate{Globe: &local, Lat: 1, Lon: 2})); v.Globe != "LOCAL 0x0003" {
		t.Fatalf("local: %q", v.Globe)
	}
	if _, err := coordinate(core.Record{Kind: core.KindValue, Codec: core.CodecCoordinateEarth}); err == nil {
		t.Fatal("broken: want error")
	}
}
