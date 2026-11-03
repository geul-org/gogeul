//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeCoordinate — the docs' Berlin (COORDINATE-EARTH, 1 arcminute, 14 WORDs) and Mars (COORDINATE-GEO, unknown, 14 WORDs) words; a 1-WORD unit gives 13 WORDs, a Geo point with a unit takes a Length extension
package core

import (
	"slices"
	"testing"
)

func TestEncodeCoordinate(t *testing.T) {
	arcminute := Record{Kind: KindReference, Codec: CodecWikidataQID, Payload: []Word{0x0003, 0x3212}}
	ws, _ := EncodeRecord(EncodeCoordinate(Coordinate{Lat: 52.516666666667, Lon: 13.383333333333, Uncertainty: 1, Unit: &arcminute}))
	if !slices.Equal(ws, []Word{0x044D, 0x404A, 0x4222, 0x2222, 0x2251, 0x402A, 0xC444, 0x4444, 0x4389, 0x3F80, 0x0000, 0x0142, 0x0003, 0x3212}) {
		t.Fatalf("earth: %04X", ws)
	}
	mars := Record{Kind: KindReference, Codec: CodecWikidataQID, Payload: []Word{111}}
	ws, _ = EncodeRecord(EncodeCoordinate(Coordinate{Globe: &mars, Lat: 4.5, Lon: 137.4}))
	if !slices.Equal(ws, []Word{0x045D, 0x0141, 0x006F, 0x4012, 0, 0, 0, 0x4061, 0x2CCC, 0xCCCC, 0xCCCD, 0, 0, Absent}) {
		t.Fatalf("geo: %04X", ws)
	}
	degree := Record{Kind: KindReference, Codec: CodecWikidataQID, Payload: []Word{28390}}
	ws, _ = EncodeRecord(EncodeCoordinate(Coordinate{Lat: 1, Lon: 2, Uncertainty: 0.0001, Unit: &degree}))
	if len(ws) != 13 || ws[0] != 0x044C || ws[9] != 0x38D1 || ws[10] != 0xB717 {
		t.Fatalf("earth, 0.0001 degree: %04X", ws)
	}
	ws, _ = EncodeRecord(EncodeCoordinate(Coordinate{Globe: &mars, Lat: 1, Lon: 2, Uncertainty: 2.5, Unit: &degree}))
	if len(ws) != 16 || ws[0] != 0x045E || ws[1] != 16 || ws[12] != 0x4020 || ws[13] != 0x0000 {
		t.Fatalf("geo with a unit: %04X", ws)
	}
}
