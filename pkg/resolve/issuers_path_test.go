//ff:func feature=resolve type=helper control=sequence
//ff:what issuersPath — test helper: the stored address of version 1 of the Issuer list (issuers/0x00000001.geul)
package resolve

// issuersPath is where version 1 of the Issuer list is stored.
func issuersPath() string {
	return versionPath(issuersKey(), 1)
}
