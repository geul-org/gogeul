//ff:func feature=resolve type=helper control=sequence
//ff:what sampleKey — test fixture: the Key of sample document id at version v
package resolve

import (
	"github.com/geul-org/gogeul/pkg/core"
)

func sampleKey(id core.Word, v uint32) Key {
	return Key{Issuer: []core.Word{0, 0}, Doc: []core.Word{0xffff, 0, 0, id}, Version: v}
}
