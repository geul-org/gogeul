//ff:func feature=notation type=encoder control=selection
//ff:what parseValueArg — the one argument of a Value call by codec: quoted bytes, a decimal, an integer of any size, a date, seconds or digest words
package notation

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/geul-org/gogeul/pkg/core"
)

// parseValueArg encodes arg as a Value of codec c.
func parseValueArg(c uint16, arg string) (core.Record, error) {
	switch even := c &^ 1; {
	case even == core.CodecUTF8TextEven || even == codecURIEven || even == core.CodecEDTFEven || even == core.CodecGEDCOMDateEven || even == core.CodecLanguageBCP47Even:
		s, err := strconv.Unquote(arg)
		return core.EncodeBytesValue(even, []byte(s)), err
	case c == core.CodecSignedDecimal || c == core.CodecUnsignedDecimal:
		return parseDecimalValue(c, arg)
	case c == core.CodecSignedInteger || c == core.CodecUnsignedInteger:
		n, ok := new(big.Int).SetString(arg, 10)
		if !ok {
			return core.Record{}, fmt.Errorf("geul/notation: %q is not an integer", arg)
		}
		return core.EncodeIntegerBig(c == core.CodecSignedInteger, n)
	case c == core.CodecGregorianDate:
		return parseDate(arg)
	case c == core.CodecPosixSeconds:
		n, err := strconv.ParseInt(arg, 10, 64)
		return core.EncodePosixSeconds(n), err
	case c == core.CodecSHA256Digest || c == core.CodecSHA1Digest:
		ws, err := parseHex(arg)
		return core.Record{Kind: core.KindValue, Codec: c, Payload: ws}, err
	default:
		return core.Record{}, fmt.Errorf("geul/notation: %s has no call form; write it as raw words", core.CodecName(core.KindValue, c))
	}
}
