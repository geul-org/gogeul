//ff:func feature=resolve type=helper control=sequence
//ff:what namesPath — test helper: the stored address of version 1 of an Issuer's document list (<Issuer>/docs/0x00000001.geul)
package resolve

// namesPath is where version 1 of issuer's document list is stored.
func namesPath(issuer string) string {
	return versionPath(namesKey(issuer), 1)
}
