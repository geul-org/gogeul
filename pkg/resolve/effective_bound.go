//ff:func feature=resolve type=decoder control=iteration dimension=1
//ff:what effectiveBound — the instant of a document's EFFECTIVE-FROM or EFFECTIVE-UNTIL Meta, if present
package resolve

import (
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// effectiveBound returns the instant of doc's Meta with the given codec
// (core.MetaEffectiveFrom or core.MetaEffectiveUntil) and whether it is present.
func effectiveBound(doc inspect.Document, codec uint16) (timeValue, bool, error) {
	for _, e := range doc.Head {
		if e.Record.Kind != core.KindMeta || e.Record.Codec != codec {
			continue
		}
		t, err := metaTime(e.Record)
		return t, err == nil, err
	}
	return timeValue{}, false, nil
}
