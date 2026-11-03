//ff:func feature=schema type=rule control=selection
//ff:what singleKey — the rule and the counting key of a HEAD record that may appear once: its Meta code, or MESSAGE and its language
package schema

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// singleKey returns the rule reporting a repeat of r and the key repeats
// share, or false when r may repeat: a HEAD Edge (SOURCE, QUALIFIER) never
// counts, a MESSAGE counts per language (messageLanguage, grammar 0x1D9L
// "언어마다 많아야 하나"), the singleRule Meta per code.
func singleKey(r rec) (string, string, bool) {
	rule, single := singleRule[r.R.Codec]
	switch {
	case r.R.Kind != core.KindMeta:
		return "", "", false
	case r.R.Codec == core.MetaMessage:
		lang, ok := messageLanguage(r.R)
		return "M8", "MESSAGE " + lang, ok
	default:
		return rule, fmt.Sprintf("0x%02X", r.R.Codec), single
	}
}
