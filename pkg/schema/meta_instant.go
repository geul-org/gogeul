//ff:func feature=schema type=decoder control=iteration dimension=1
//ff:what metaInstant — the instant of the first HEAD Meta with a codec, when it carries a valid time Value
package schema

import (
	"time"

	"github.com/geul-org/gogeul/pkg/core"
)

// metaInstant returns the instant of the first Meta of codec c in head.
func metaInstant(head []rec, c uint16) (time.Time, rec, bool) {
	for _, r := range head {
		if r.R.Codec != c {
			continue
		}
		v, ok := embeddedValue(r)
		t, err := core.DecodeTimeValue(v.R)
		return t, r, ok && err == nil
	}
	return time.Time{}, rec{}, false
}
