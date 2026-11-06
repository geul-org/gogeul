//ff:type feature=core type=model
//ff:what Value codec constants — the direct Value codes (Codec = code - 0x60, sections 3·4) in table order
package core

// Direct Value codes (sections 3·4, TAG 0x060L~0x081L; the Record Codec is
// code - 0x60). The codes run in table
// order without gaps and carry no family bits: a decoder that does not know
// a code does not know its family. Variable-length byte strings come in
// even/odd pairs that start on an even C: the low bit is 1 when the byte count
// is odd and the last word carries a 0x00 pad byte.
const (
	// CodecBool is BOOL (0x060L).
	CodecBool uint16 = 0x00
	// CodecSignedInteger is SIGNED-INTEGER (0x061L).
	CodecSignedInteger uint16 = 0x01
	// CodecUnsignedInteger is UNSIGNED-INTEGER (0x062L).
	CodecUnsignedInteger uint16 = 0x02
	// CodecBinaryFloat is BINARY-FLOAT (0x063L).
	CodecBinaryFloat uint16 = 0x03
	// CodecCoordinateEarth is COORDINATE-EARTH (0x064L).
	CodecCoordinateEarth uint16 = 0x04
	// CodecCoordinateGeo is COORDINATE-GEO (0x065L).
	CodecCoordinateGeo uint16 = 0x05
	// CodecQuantity is QUANTITY (0x066L).
	CodecQuantity uint16 = 0x06
	// CodecFraction is FRACTION (0x067L).
	CodecFraction uint16 = 0x07
	// CodecSignedDecimal is SIGNED-DECIMAL (0x068L).
	CodecSignedDecimal uint16 = 0x08
	// CodecUnsignedDecimal is UNSIGNED-DECIMAL (0x069L).
	CodecUnsignedDecimal uint16 = 0x09
	// CodecGregorianDate is GREGORIAN-DATE (0x06AL).
	CodecGregorianDate uint16 = 0x0A
	// CodecCivilTime is CIVIL-TIME-NANOSECONDS (0x06BL).
	CodecCivilTime uint16 = 0x0B
	// CodecPosixSeconds is POSIX-TIMESTAMP-SECONDS (0x06CL).
	CodecPosixSeconds uint16 = 0x0C
	// CodecDurationSeconds is DURATION-SECONDS (0x06DL).
	CodecDurationSeconds uint16 = 0x0D
	// CodecEDTFEven is EDTF, even byte count (0x06EL; odd 0x06FL).
	CodecEDTFEven uint16 = 0x0E
	// CodecGEDCOMDateEven is GEDCOM-DATE, even byte count (0x070L; odd 0x071L).
	CodecGEDCOMDateEven uint16 = 0x10
	// CodecUTF8TextEven is UTF8-TEXT, even byte count (0x072L; odd 0x073L).
	CodecUTF8TextEven uint16 = 0x12
	// CodecURIEven is URI, even byte count (0x074L; odd 0x075L).
	CodecURIEven uint16 = 0x14
	// CodecLanguageBCP47Even is LANGUAGE-BCP47-CODE, even byte count (0x076L; odd 0x077L).
	CodecLanguageBCP47Even uint16 = 0x16
	// CodecUTF8LanguageTextEven is UTF8-LANGUAGE-TEXT, even total byte count (0x078L; odd 0x079L).
	CodecUTF8LanguageTextEven uint16 = 0x18
	// CodecLanguageShortCode is LANGUAGE-SHORT-CODE (0x07AL): one word, a code
	// of the language registry.
	CodecLanguageShortCode uint16 = 0x1A
	// CodecUUID is UUID (0x07BL).
	CodecUUID uint16 = 0x1B
	// CodecSHA256Digest is SHA-256-DIGEST (0x07CL).
	CodecSHA256Digest uint16 = 0x1C
	// CodecSHA1Digest is SHA1-DIGEST (0x07DL).
	CodecSHA1Digest uint16 = 0x1D
	// CodecOpaqueBytesEven is OPAQUE-BYTES, even byte count (0x07EL; odd
	// 0x07FL). Formerly the GEUL Issuer's ISSUER formats 0xC000/0xC001.
	CodecOpaqueBytesEven uint16 = 0x1E
	// CodecBitVector is BIT-VECTOR (0x080L): a 2-word bit count N, then
	// ceil(N/16) words MSB-first. Formerly GEUL ISSUER format 0xC002 (§3.14).
	CodecBitVector uint16 = 0x20
	// CodecDurationNanoseconds is DURATION-NANOSECONDS (0x081L): signed int64
	// nanoseconds in 4 words. Formerly GEUL ISSUER format 0x4000 (§3.14).
	CodecDurationNanoseconds uint16 = 0x21
)
