//ff:func feature=schema type=rule control=selection
//ff:what checkCoordinate — V8: a COORDINATE-EARTH or COORDINATE-GEO has its layout, a globe in the closed list (Geo), finite degrees in range, a finite binary32 uncertainty ≥ +0, an Identified unit Reference, or ABSENT with uncertainty +0
package schema

import (
	"fmt"
	"math"

	"github.com/geul-org/gogeul/pkg/core"
)

// checkCoordinate applies V8 to a coordinate Value: COORDINATE-EARTH is 10
// WORDs then the unit; COORDINATE-GEO starts with one whole Reference record
// naming the body, from the closed globe list (globeCodec: an Identified 5Q
// code other than WIKIDATA-PID, grammar 0x065L). Latitude and longitude are finite,
// latitude in [-90, 90], longitude in [-180, 180] on Earth and [-360, 360] on
// another body (the Wikidata globecoordinate range). The uncertainty (binary32) is finite and ≥ 0 written as +0
// (not -0). The unit is a Reference of Quantifier Identified, or ABSENT
// (uncertainty unknown) with the uncertainty +0.
func checkCoordinate(r rec) []Fact {
	c, err := core.DecodeCoordinate(r.R)
	lonMax := 180.0
	if r.R.Codec == core.CodecCoordinateGeo {
		lonMax = 360
	}
	lonMin := -lonMax
	switch {
	case err != nil:
		return []Fact{factAt("V8", Error, r, "Payload", "Globe Reference(Geo), 10 WORD, Unit Reference 또는 ABSENT", err.Error())}
	case c.Globe != nil && !globeCodec(c.Globe.Codec):
		return []Fact{factAt("V8", Error, r, "Payload", "Globe: 5Q 체계(WIKIDATA-PID 제외)의 Identified", fmt.Sprintf("0x%02X", c.Globe.Codec))}
	case math.IsNaN(c.Lat) || math.IsNaN(c.Lon) || math.IsInf(c.Lat, 0) || math.IsInf(c.Lon, 0):
		return []Fact{factAt("V8", Error, r, "Payload", "유한한 위도·경도", fmt.Sprintf("%v, %v", c.Lat, c.Lon))}
	case c.Lat < -90 || c.Lat > 90 || c.Lon < lonMin || c.Lon > lonMax:
		return []Fact{factAt("V8", Error, r, "Payload", fmt.Sprintf("위도 [-90, 90], 경도 [%v, %v]", lonMin, lonMax), fmt.Sprintf("%v, %v", c.Lat, c.Lon))}
	case math.IsNaN(float64(c.Uncertainty)) || math.IsInf(float64(c.Uncertainty), 0) || math.Signbit(float64(c.Uncertainty)):
		return []Fact{factAt("V8", Error, r, "Uncertainty", "유한한 +0 이상 binary32", fmt.Sprintf("%v", c.Uncertainty))}
	case c.Unit != nil && notIdentified(c.Unit.Codec):
		return []Fact{factAt("V8", Error, r, "Unit", "Unit Quantifier Identified", fmt.Sprintf("0x%02X", c.Unit.Codec))}
	case c.Unit == nil && c.Uncertainty != 0:
		return []Fact{factAt("V8", Error, r, "Uncertainty", "Unit이 ABSENT(미상)이면 +0", fmt.Sprintf("%v", c.Uncertainty))}
	default:
		return nil
	}
}
