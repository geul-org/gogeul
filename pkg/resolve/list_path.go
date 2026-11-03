//ff:func feature=resolve type=helper control=sequence
//ff:what listPath — `<Issuer>/docs/<문서 ID>/versions.geul` (codecs/<코덱 ID>/ for a codec), the version list address
package resolve

// listPath is the address of d's version list document, relative to
// `.well-known/geul/` (docFolder + versions.geul).
func listPath(d DocKey) string {
	return docFolder(d) + "/" + listFile
}
