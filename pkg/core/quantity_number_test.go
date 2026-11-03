//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestQuantityNumber — the six numeric codecs pass; BOOL, QUANTITY, a coordinate and a Reference do not
package core

import "testing"

func TestQuantityNumber(t *testing.T) {
	for _, c := range []uint16{CodecSignedInteger, CodecUnsignedInteger, CodecBinaryFloat, CodecFraction, CodecSignedDecimal, CodecUnsignedDecimal} {
		if !quantityNumber(Record{Kind: KindValue, Codec: c}) {
			t.Fatalf("0x%02X: want true", c)
		}
	}
	for _, r := range []Record{{Kind: KindValue, Codec: 0x00}, {Kind: KindValue, Codec: CodecQuantity}, {Kind: KindValue, Codec: CodecCoordinateEarth}, {Kind: KindReference, Codec: CodecSignedInteger}} {
		if quantityNumber(r) {
			t.Fatalf("%+v: want false", r)
		}
	}
}
