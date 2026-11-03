//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what Dir.CodecSpec — the codec specs of every direct codec spec document published under the GEUL Issuer
package resolve

import (
	"github.com/geul-org/gogeul/pkg/codecspec"
)

// CodecSpec gathers the specs of the latest version of every codec spec
// document, the GEUL Issuer's codecs `0x00000000/codecs/<codecDocID>/` that
// have a version list (GEUL-ORG 직접 코드 코덱북), in ID order. Other codecs
// of the Issuer are skipped. The first error stops it.
func (d Dir) CodecSpec() (codecspec.Spec, error) {
	is := geulIssuer + "/" + codecsDir
	ents, _ := d.store().ReadDir(is)
	s := codecspec.Spec{}
	for _, e := range ents {
		target, ok := codecTargetOf(e.Name())
		if _, err := d.store().ReadFile(is + "/" + e.Name() + "/" + listFile); !ok || err != nil {
			continue
		}
		spec, err := d.codecSpecOf(target)
		if err != nil {
			return nil, err
		}
		s = s.Merge(spec)
	}
	return s, nil
}
