//ff:func feature=core type=decoder control=sequence
//ff:what DecodeBytesValue — the bytes of an even/odd byte-string Value record, enforcing the pad byte
package core

import "fmt"

// DecodeBytesValue returns the bytes of r, which must be a Value record of the
// even/odd codec pair whose even member is evenCodec. The byte count is
// 2 × payload words − the codec's low bit; an odd record's pad byte must be 0x00
// and an odd record must have at least one payload word.
func DecodeBytesValue(r Record, evenCodec uint16) ([]byte, error) {
	if r.Kind != KindValue || r.Codec&^1 != evenCodec {
		return nil, fmt.Errorf("geul/core: expected Value codec 0x%02X/0x%02X, got kind %d codec 0x%02X", evenCodec, evenCodec|1, r.Kind, r.Codec)
	}
	b := PutWords(r.Payload)
	if r.Codec&1 == 0 {
		return b, nil
	}
	if len(b) == 0 || b[len(b)-1] != 0x00 {
		return nil, fmt.Errorf("geul/core: odd byte-string Value codec 0x%02X needs a 0x00 pad byte", r.Codec)
	}
	return b[:len(b)-1], nil
}
