//ff:func feature=core type=encoder control=sequence
//ff:what EncodeBytesValue — a byte-string Value record using the even/odd codec pair
package core

// EncodeBytesValue builds the Value record for b under the even/odd codec pair
// whose even member is evenCodec (GEUL t002 OVERVIEW Value): the codec's low bit
// is the byte-count parity and an odd count is padded with one 0x00 byte. The
// bytes are stored verbatim (no normalization).
func EncodeBytesValue(evenCodec uint16, b []byte) Record {
	odd := uint16(len(b) % 2)
	padded := b
	if odd == 1 {
		padded = append(append([]byte(nil), b...), 0x00)
	}
	data, _ := ReadWords(padded)
	return Record{Kind: KindValue, Codec: evenCodec | odd, Payload: data}
}
