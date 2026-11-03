//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodeCoordinate — Earth and Geo round trip with a binary32 uncertainty and unit, or ABSENT; another codec, a missing globe, a short payload, a unit that is not a Reference or ABSENT, or words after it fail
package core

import "testing"

func TestDecodeCoordinate(t *testing.T) {
	second := Record{Kind: KindReference, Codec: CodecWikidataQID, Payload: []Word{0x000C, 0xA691}}
	in := Coordinate{Lat: -33.5, Lon: 151.25, Uncertainty: 0.1, Unit: &second}
	if c, err := DecodeCoordinate(EncodeCoordinate(in)); err != nil || c.Lat != in.Lat || c.Lon != in.Lon || c.Uncertainty != 0.1 || c.Unit == nil || c.Unit.Payload[1] != 0xA691 || c.Globe != nil {
		t.Fatalf("earth: %+v %v", c, err)
	}
	moon := Record{Kind: KindReference, Codec: CodecWikidataQID, Payload: []Word{405}}
	if c, err := DecodeCoordinate(EncodeCoordinate(Coordinate{Globe: &moon, Lat: 1, Lon: 359.5})); err != nil || c.Globe.Payload[0] != 405 || c.Lon != 359.5 || c.Unit != nil {
		t.Fatalf("geo: %+v %v", c, err)
	}
	unknownUnit := EncodeCoordinate(Coordinate{})
	unknownUnit.Payload[10] = UnknownValue
	trailing := EncodeCoordinate(Coordinate{})
	trailing.Payload = append(trailing.Payload, 0)
	value := EncodeCoordinate(Coordinate{})
	value.Payload = append(value.Payload[:10], 0x0882, 5)
	for i, r := range []Record{
		EncodeUintValue(1),
		{Kind: KindValue, Codec: CodecCoordinateGeo, Payload: make([]Word, 11)},
		{Kind: KindValue, Codec: CodecCoordinateGeo, Payload: []Word{0x08C0}},
		{Kind: KindValue, Codec: CodecCoordinateEarth, Payload: make([]Word, 10)},
		unknownUnit, trailing, value,
	} {
		if _, err := DecodeCoordinate(r); err == nil {
			t.Fatalf("case %d: want error", i)
		}
	}
}
