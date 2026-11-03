//ff:func feature=resolve type=helper control=sequence
//ff:what addrDoc — `<문서 ID>/<tail>` under an Issuer's docs/: a 16-digit document ID other than the reserved Issuer list ID, then the version file
package resolve

import "github.com/geul-org/gogeul/pkg/core"

// addrDoc reads `<문서 ID>/<name>`: the 4-word document ID (16 hex
// digits), the name as addrTail. The reserved Issuer list ID
// (core.DocIDIssuers) is not a document address.
func addrDoc(a Addr, id, name string, bad error) (Addr, error) {
	b, ok := parseHexID(id, 4)
	if !ok {
		return Addr{}, bad
	}
	doc := hexWords(b)
	if wordsUint(doc) == core.DocIDIssuers {
		return Addr{}, bad
	}
	a.Kind, a.ID = AddrDoc, doc
	return addrTail(a, name, bad)
}
