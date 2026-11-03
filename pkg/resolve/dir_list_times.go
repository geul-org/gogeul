//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what Dir.listTimes — the record count and RECORDED-AT of a version list and of each earlier list it reaches by PARENT
package resolve

import (
	"context"

	"github.com/geul-org/gogeul/pkg/core"
)

// listTimes walks the version list b and its earlier lists by PARENT through
// the Issuer's `hashes/`, newest first, for at most limit lists. It stops at a list that is
// missing or does not split.
func (d Dir) listTimes(issuer string, b []byte, limit int) []listTime {
	var out []listTime
	for i := 0; i < limit && b != nil; i++ {
		ws, err := core.ReadWords(b)
		if err != nil {
			break
		}
		doc, err := core.SplitDocument(ws)
		if err != nil {
			break
		}
		out = append(out, listTime{n: len(doc.Body), at: recordedAt(doc.Head)})
		parent, ok := listParent(doc.Head)
		if !ok {
			break
		}
		b, _ = d.Fetch(context.Background(), hashPath(issuer, parent))
	}
	return out
}
