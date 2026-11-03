//ff:func feature=notation type=encoder control=selection
//ff:what parseDecimalValue — a SIGNED-DECIMAL (any size) or UNSIGNED-DECIMAL (int64 coefficient) from decimal text
package notation

import (
	"github.com/geul-org/gogeul/pkg/core"
)

// parseDecimalValue reads arg (parseDecimal) as codec c.
func parseDecimalValue(c uint16, arg string) (core.Record, error) {
	e, n, err := parseDecimal(arg)
	switch {
	case err != nil:
		return core.Record{}, err
	case c == core.CodecSignedDecimal:
		return core.EncodeSignedDecimalBig(e, n), nil
	case !n.IsInt64():
		return core.Record{}, errTooBig(arg)
	default:
		return core.EncodeDecimalValue(false, e, n.Int64())
	}
}
