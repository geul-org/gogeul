//ff:func feature=resolve type=helper control=sequence
//ff:what issuersKey — the DocKey of geul.org's Issuer list (BIND names, the old issuers.geul): 0x00000000 and the reserved ID 0xffffffffffffffff
package resolve

// issuersKey is the DocKey of the Issuer list (`issuers/`, GEUL-ORG Issuer
// 이름), a versioned document of geul.org under the reserved ID
// core.DocIDIssuers.
func issuersKey() DocKey {
	return DocKey{Issuer: geulIssuer, Doc: listIDIssuers}
}
