//ff:type feature=core type=model
//ff:what Value codec constants — the direct Value codes (C, KKK 001) in table order
package core

// Direct Value codes (KKK 001, TAG 0x040L~0x05DL). The codes run in table
// order without gaps and carry no family bits: a decoder that does not know
// a code does not know its family. Variable-length byte strings come in
// even/odd pairs that start on an even C: the low bit is 1 when the byte count
// is odd and the last word carries a 0x00 pad byte.
const (
	// CodecBool is BOOL (0x040L).
	CodecBool uint16 = 0x00
	// CodecSignedInteger is SIGNED-INTEGER (0x041L).
	CodecSignedInteger uint16 = 0x01
	// CodecUnsignedInteger is UNSIGNED-INTEGER (0x042L).
	CodecUnsignedInteger uint16 = 0x02
	// CodecBinaryFloat is BINARY-FLOAT (0x043L).
	CodecBinaryFloat uint16 = 0x03
	// CodecCoordinateEarth is COORDINATE-EARTH (0x044L).
	CodecCoordinateEarth uint16 = 0x04
	// CodecCoordinateGeo is COORDINATE-GEO (0x045L).
	CodecCoordinateGeo uint16 = 0x05
	// CodecQuantity is QUANTITY (0x046L).
	CodecQuantity uint16 = 0x06
	// CodecFraction is FRACTION (0x047L).
	CodecFraction uint16 = 0x07
	// CodecSignedDecimal is SIGNED-DECIMAL (0x048L).
	CodecSignedDecimal uint16 = 0x08
	// CodecUnsignedDecimal is UNSIGNED-DECIMAL (0x049L).
	CodecUnsignedDecimal uint16 = 0x09
	// CodecGregorianDate is GREGORIAN-DATE (0x04AL).
	CodecGregorianDate uint16 = 0x0A
	// CodecCivilTime is CIVIL-TIME-NANOSECONDS (0x04BL).
	CodecCivilTime uint16 = 0x0B
	// CodecPosixSeconds is POSIX-TIMESTAMP-SECONDS (0x04CL).
	CodecPosixSeconds uint16 = 0x0C
	// CodecDurationSeconds is DURATION-SECONDS (0x04DL).
	CodecDurationSeconds uint16 = 0x0D
	// CodecEDTFEven is EDTF, even byte count (0x04EL; odd 0x04FL).
	CodecEDTFEven uint16 = 0x0E
	// CodecGEDCOMDateEven is GEDCOM-DATE, even byte count (0x050L; odd 0x051L).
	CodecGEDCOMDateEven uint16 = 0x10
	// CodecUTF8TextEven is UTF8-TEXT, even byte count (0x052L; odd 0x053L).
	CodecUTF8TextEven uint16 = 0x12
	// CodecURIEven is URI, even byte count (0x054L; odd 0x055L).
	CodecURIEven uint16 = 0x14
	// CodecLanguageBCP47Even is LANGUAGE-BCP47-CODE, even byte count (0x056L; odd 0x057L).
	CodecLanguageBCP47Even uint16 = 0x16
	// CodecUTF8LanguageTextEven is UTF8-LANGUAGE-TEXT, even total byte count (0x058L; odd 0x059L).
	CodecUTF8LanguageTextEven uint16 = 0x18
	// CodecLanguageShortCode is LANGUAGE-SHORT-CODE (0x05AL): one word, a code
	// of the language registry.
	CodecLanguageShortCode uint16 = 0x1A
	// CodecUUID is UUID (0x05BL).
	CodecUUID uint16 = 0x1B
	// CodecSHA256Digest is SHA-256-DIGEST (0x05CL).
	CodecSHA256Digest uint16 = 0x1C
	// CodecSHA1Digest is SHA1-DIGEST (0x05DL).
	CodecSHA1Digest uint16 = 0x1D
)
