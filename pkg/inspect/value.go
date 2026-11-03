//ff:type feature=inspect type=model
//ff:what Value — a decoded Value record of a codec this library knows
package inspect

import "time"

// Value is a decoded Value record. Codec is the codec as written (the odd
// member of a byte-string pair kept). Text holds UTF8-TEXT, BCP 47 tags, the
// text of a UTF8-LANGUAGE-TEXT (its language in Lang: the tag; a short code
// in hex, or its registry tag when a codec spec names it, namedLanguage) and a SHA-256-DIGEST as `0x` hex, Time a GREGORIAN-DATE or
// POSIX-TIMESTAMP-SECONDS (UTC), Uint an UNSIGNED-INTEGER, and Exp/Coef a
// SIGNED- or UNSIGNED-DECIMAL as written (Coef×10^Exp, quantum kept), Float a
// BINARY-FLOAT (binaryFloat), and
// Lat/Lon/Float/Unit a COORDINATE-EARTH or COORDINATE-GEO, whose body is in
// Globe (referenceText; empty for Earth) and whose uncertainty (binary32,
// widened) is Float in the unit Unit (referenceText; empty when ABSENT, the uncertainty unknown),
// Amount/Unit/Lower/Upper a QUANTITY (exact numbers and the unit as
// text; empty when left out), and Number a SIGNED-INTEGER (digits, any size)
// or a FRACTION (n/d as written, not reduced).
type Value struct {
	Codec  uint16
	Text   string
	Lang   string
	Time   time.Time
	Uint   uint64
	Exp    int32
	Coef   int64
	Float  float64
	Lat    float64
	Lon    float64
	Globe  string
	Amount string
	Unit   string
	Lower  string
	Upper  string
	Number string
}
