//ff:func feature=resolve type=helper control=sequence
//ff:what addrDoc — `<문서 ID>/<tail>` under an Issuer's docs/: a 16-digit document ID (or a name via docOf), then the version file
package resolve

import "github.com/geul-org/gogeul/pkg/core"

// addrDoc reads `<문서 ID>/<name>`: docOf gives the 4-word document ID (16
// lowercase hex digits; Dir.Find also takes a document name), the name as
// addrTail.
func addrDoc(a Addr, id, name string, docOf func([]core.Word, string) ([]core.Word, error), bad error) (Addr, error) {
	doc, err := docOf(a.Issuer, id)
	if err != nil {
		return Addr{}, err
	}
	if len(doc) != 4 {
		return Addr{}, bad
	}
	a.Kind, a.ID = AddrDoc, doc
	return addrTail(a, name, bad)
}
