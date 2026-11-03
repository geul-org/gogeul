//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what recordedAt — the RECORDED-AT time of a version list HEAD; zero when absent or unreadable
package resolve

import (
	"time"

	"github.com/geul-org/gogeul/pkg/core"
)

// recordedAt reads the RECORDED-AT Meta of head.
func recordedAt(head []core.Record) time.Time {
	for _, r := range head {
		if r.Kind != core.KindMeta || r.Codec != core.MetaRecordedAt {
			continue
		}
		tv, _ := metaTime(r)
		return tv.Time
	}
	return time.Time{}
}
