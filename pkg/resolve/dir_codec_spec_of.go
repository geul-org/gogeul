//ff:func feature=resolve type=resolver control=sequence
//ff:what Dir.codecSpecOf — the codec specs one codec spec document's latest version declares; a DECLARE of another Kind than the target is an error
package resolve

import (
	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// codecSpecOf collects the BODY of target's latest codec spec document, which
// must declare a codec of target's Kind (checkDeclareKind, §1.1 103).
func (d Dir) codecSpecOf(target core.Word) (codecspec.Spec, error) {
	b, err := d.latestCodec(target)
	if err != nil {
		return nil, err
	}
	doc, _ := inspect.Decode(b)
	if err := checkDeclareKind(doc, core.KindOf(target)); err != nil {
		return nil, err
	}
	return specOfDoc(doc), nil
}
