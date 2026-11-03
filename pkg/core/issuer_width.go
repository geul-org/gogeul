//ff:func feature=core type=helper control=sequence
//ff:what IssuerWidth — the Issuer word count of an ISSUER codec, 0 for a direct codec
package core

// IssuerWidth returns how many Issuer words the extension of a k/codec record
// carries: 2, 4 or 16 for the last three codes of the Reference, Value and Edge
// zones, read from the low 2 bits of C (01, 10, 11), and 0 for every other
// codec. Meta has no ISSUER codes. The code before the three (the empty
// 0x05CL, 0x09CL and 0x1BCL) is not an ISSUER.
func IssuerWidth(k Kind, codec uint16) int {
	if k == KindMeta || codec <= CodecMax(k)-3 || codec > CodecMax(k) {
		return 0
	}
	return [...]int{0, 2, 4, 16}[codec&3]
}
