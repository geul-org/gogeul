//ff:func feature=inspect type=decoder control=selection
//ff:what ParseValue — a Value record of a known codec as Value; false for any other codec or a malformed payload
package inspect

import "github.com/geul-org/gogeul/pkg/core"

// ParseValue decodes the Value codecs this library knows: the byte strings
// (UTF8-TEXT, URI, LANGUAGE-BCP47-CODE, EDTF, GEDCOM-DATE; either pair member)
// as Text, UTF8-LANGUAGE-TEXT as Text and Lang (languageText), the two coordinates
// (coordinate), BINARY-FLOAT up to binary64 (binaryFloat), UNSIGNED-INTEGER, the two decimals, the two time
// codecs, QUANTITY (quantity), SIGNED-INTEGER and FRACTION as Number
// (numberText), SHA-256-DIGEST and SHA1-DIGEST (in hex). ok is
// false for any other codec or a malformed payload.
func ParseValue(r core.Record) (Value, bool) {
	v := Value{Codec: r.Codec}
	var err error
	if r.Codec == core.CodecBinaryFloat {
		v.Float, err = binaryFloat(r.Payload)
		return v, err == nil
	}
	if r.Codec&^1 == core.CodecCoordinateEarth {
		v, err = coordinate(r)
		return v, err == nil
	}
	if r.Codec == core.CodecQuantity {
		v.Amount, v.Unit, v.Lower, v.Upper, err = quantity(r)
		return v, err == nil
	}
	if r.Codec == core.CodecSignedInteger || r.Codec == core.CodecFraction {
		v.Number, err = numberText(r)
		return v, err == nil
	}
	if r.Codec == core.CodecSHA256Digest || r.Codec == core.CodecSHA1Digest {
		v.Text = core.FormatID(r.Payload)
		return v, len(r.Payload) == map[uint16]int{core.CodecSHA256Digest: 16, core.CodecSHA1Digest: 10}[r.Codec]
	}
	switch even := r.Codec &^ 1; even {
	case core.CodecUTF8TextEven, core.CodecURIEven, core.CodecLanguageBCP47Even, core.CodecEDTFEven, core.CodecGEDCOMDateEven:
		var b []byte
		b, err = core.DecodeBytesValue(r, even)
		v.Text = string(b)
	case core.CodecUTF8LanguageTextEven:
		v.Text, v.Lang, err = languageText(r)
	case core.CodecUnsignedInteger:
		v.Uint, err = core.DecodeUintValue(r)
	case core.CodecSignedDecimal:
		v.Exp, v.Coef, err = core.DecodeDecimalValue(r)
	case core.CodecGregorianDate, core.CodecPosixSeconds:
		v.Time, err = core.DecodeTimeValue(r)
	default:
		return Value{}, false
	}
	return v, err == nil
}
