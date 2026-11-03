//ff:func feature=resolve type=decoder control=iteration dimension=1
//ff:what declaredTarget — the target TAG of the one DECLARE in a document's BODY (direct or ISSUER target); false for none, several or unreadable bytes
package resolve

import "github.com/geul-org/gogeul/pkg/core"

// declaredTarget reads document b and returns the target word of its one
// DECLARE, whatever the target (a direct code or an ISSUER form).
func declaredTarget(b []byte) (core.Word, bool) {
	ws, err := core.ReadWords(b)
	if err != nil {
		return 0, false
	}
	doc, err := core.SplitDocument(ws)
	if err != nil {
		return 0, false
	}
	var found []core.Word
	for _, r := range doc.Body {
		if dc, err := core.DecodeDeclare(r); err == nil {
			found = append(found, dc.Target)
		}
	}
	if len(found) != 1 {
		return 0, false
	}
	return found[0], true
}
