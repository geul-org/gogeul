//ff:func feature=schema type=rule control=selection
//ff:what checkValue — V2~V10 for one Value record; an unassigned direct code passes
package schema

import "github.com/geul-org/gogeul/pkg/core"

// checkValue applies the Value rules to a direct-code Value. ISSUER formats are
// left to layer 2 (F5 already checked their extension TAG). A code with no
// rule here (unassigned) is not reported: the reserved-code Warn V1 was
// removed; an unassigned code shows in codec lookup.
func checkValue(r rec) []Fact {
	c := r.R.Codec
	switch {
	case c >= core.CodecValueIssuer2Word:
		return nil
	case c == core.CodecBool:
		return checkBool(r)
	case c == core.CodecSignedInteger || c == core.CodecUnsignedInteger:
		return checkInteger(r, r.R.Payload, c == core.CodecSignedInteger, "Payload")
	case c == core.CodecCoordinateEarth || c == core.CodecCoordinateGeo:
		return checkCoordinate(r)
	case c == core.CodecQuantity:
		return checkQuantity(r)
	case c == core.CodecFraction:
		return checkFraction(r)
	case c == core.CodecBinaryFloat:
		return checkWidth(r, 1, 2, 4, 8)
	case c == core.CodecSignedDecimal || c == core.CodecUnsignedDecimal:
		return checkDecimal(r, c == core.CodecSignedDecimal)
	case c == core.CodecCivilTime:
		return checkCivilTime(r)
	case c == core.CodecLanguageShortCode:
		return checkShortCode(r)
	case fixedValueWidth[c] > 0:
		return checkWidth(r, fixedValueWidth[c])
	case c == core.CodecUTF8TextEven, c == core.CodecUTF8TextEven|1, c >= core.CodecURIEven && c <= core.CodecLanguageBCP47Even|1, c == core.CodecEDTFEven, c == core.CodecEDTFEven|1,
		c == core.CodecGEDCOMDateEven, c == core.CodecGEDCOMDateEven|1:
		return checkBytes(r)
	case c == core.CodecUTF8LanguageTextEven, c == core.CodecUTF8LanguageTextEven|1:
		return checkLanguageText(r)
	case c == core.CodecOpaqueBytesEven, c == core.CodecOpaqueBytesEven|1:
		return checkBytes(r)
	case c == core.CodecBitVector:
		return checkBitVector(r)
	default:
		return nil
	}
}
