//ff:func feature=resolve type=helper control=sequence
//ff:what namesKey — the DocKey of an Issuer's document list (BIND names, the old documents.geul): the reserved ID 0xfffffffffffffffe
package resolve

// namesKey is the DocKey of issuer's document list (`<Issuer>/docs/`, GEUL-ORG
// 문서 이름), a versioned document under the reserved ID core.DocIDDocs.
func namesKey(issuer string) DocKey {
	return DocKey{Issuer: issuer, Doc: listIDDocs}
}
