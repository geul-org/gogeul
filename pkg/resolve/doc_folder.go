//ff:func feature=resolve type=helper control=sequence
//ff:what docFolder — `<Issuer>/docs/<문서 ID>` for a 4-word document ID, `<Issuer>/codecs/<코덱 ID>` for a 2-word codec ID
package resolve

// docFolder is the folder of d's versions and version list: under docs/ for a
// document (16 hex digits) and under codecs/ for a codec (8 hex digits, a
// codec document names itself by CODEC-POINTER, not a document ID).
func docFolder(d DocKey) string {
	if len(d.Doc) == 2+8 {
		return d.Issuer + "/" + codecsDir + "/" + d.Doc
	}
	return d.Issuer + "/" + docsDir + "/" + d.Doc
}
