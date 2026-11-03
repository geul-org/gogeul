//ff:func feature=resolve type=helper control=sequence
//ff:what sampleKeyOf — test fixture: the DocKey of edgeDoc(id)
package resolve

import "github.com/geul-org/gogeul/pkg/core"

// sampleKeyOf is the DocKey of the document edgeDoc builds for id.
func sampleKeyOf(id core.Word) DocKey {
	return Key{Issuer: []core.Word{0, 0}, Doc: []core.Word{0xffff, 0, 0, id}}.DocKey()
}
