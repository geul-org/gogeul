//ff:func feature=resolve type=helper control=sequence
//ff:what issuersPath — where geul.org's Issuer list (BIND names, the old issuers.geul) is kept: `issuers/0x00000001.geul`
package resolve

// issuersPath is the stored address of the Issuer list (GEUL-ORG Issuer 이름;
// `issuers/` list, listDocPath).
func issuersPath() string {
	return listDocPath(issuersDir)
}
