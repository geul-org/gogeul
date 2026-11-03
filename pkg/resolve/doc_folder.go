//ff:func feature=resolve type=helper control=selection
//ff:what docFolder — the folder of a DocKey: <Issuer>/docs/<문서 ID>, <Issuer>/codecs/<코덱 ID>, or issuers for the reserved Issuer list ID
package resolve

// docFolder is the folder of d's versions and version list: under docs/ for a
// document (16 hex digits), under codecs/ for a codec (8 hex digits), and
// `issuers` for the reserved Issuer list document ID (core.DocIDIssuers).
func docFolder(d DocKey) string {
	switch {
	case len(d.Doc) == 2+8:
		return d.Issuer + "/" + codecsDir + "/" + d.Doc
	case d.Doc == listIDIssuers:
		return issuersDir
	default:
		return d.Issuer + "/" + docsDir + "/" + d.Doc
	}
}
