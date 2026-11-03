//ff:func feature=resolve type=resolver control=selection
//ff:what Dir.docOfName — a document ID written as canonical hex or as a name of the Issuer's document list
package resolve

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// docOfName returns the document ID s names under issuer: 4-word canonical
// hex as is, otherwise the name looked up in the Issuer's document list
// (`<Issuer>/docs/`); a name bound to a codec gives its 2-word codec ID.
func (d Dir) docOfName(issuer []core.Word, s string) ([]core.Word, error) {
	if b, ok := parseHexID(s, 4); ok {
		return hexWords(b), nil
	}
	_, _, names, err := d.namesOf(namesPath(core.FormatID(issuer)))
	if err != nil {
		return nil, err
	}
	k, ok := names[s]
	if !ok {
		return nil, fmt.Errorf("%w: document name %q under %s", ErrNotFound, s, core.FormatID(issuer))
	}
	b, _ := parseHexID(k.Doc, (len(k.Doc)-2)/4)
	return hexWords(b), nil
}
