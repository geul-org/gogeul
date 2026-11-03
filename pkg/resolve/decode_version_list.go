//ff:func feature=resolve type=decoder control=iteration dimension=1
//ff:what decodeVersionList — a version list document as its versions, records and content hash
package resolve

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// decodeVersionList reads the version list document of the document whose
// Issuer and ID are issuer (any accepted width, normalIssuer) and docHex in
// canonical hex (edge/PUBLICATION), with its PARENT (listParent). Anything
// that does not read as one, a record of another Issuer included, is a
// verification failure.
func decodeVersionList(b []byte, issuer, docHex string) (versionLog, error) {
	if n, ok := normalIssuer(issuer); ok {
		issuer = n
	}
	ws, err := core.ReadWords(b)
	if err != nil {
		return versionLog{}, fmt.Errorf("%w: %v", ErrVerify, err)
	}
	doc, err := core.SplitDocument(ws)
	if err != nil {
		return versionLog{}, fmt.Errorf("%w: %v", ErrVerify, err)
	}
	h, _ := core.ContentHash(ws)
	lg := versionLog{Records: doc.Body, Hash: h, Present: true}
	lg.Parent, lg.HasParent = listParent(doc.Head)
	lg.Next = listNext(doc.Head)
	for _, r := range doc.Body {
		if err := applyListRecord(&lg, r, issuer, docHex); err != nil {
			return versionLog{}, fmt.Errorf("%w: version list: %v", ErrVerify, err)
		}
	}
	return lg, nil
}
