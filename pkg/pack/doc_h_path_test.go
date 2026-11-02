//ff:func feature=pack type=helper control=sequence
//ff:what docHPath — the h/ path of a GEUL document's bytes
package pack

// docHPath names doc by its content hash.
func docHPath(doc []byte) string {
	h, _ := contentHash(doc)
	return hPath(string(h[:]))
}
