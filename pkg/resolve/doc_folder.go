//ff:func feature=resolve type=helper control=selection
//ff:what docFolder — the folder of a DocKey: <Issuer>/docs/<문서 ID>, <Issuer>/codecs/<코덱 ID>, or a list folder (issuers, <Issuer>/docs, <Issuer>/codecs) for a reserved list ID
package resolve

// docFolder is the folder of d's versions and version list: under docs/ for a
// document (16 hex digits), under codecs/ for a codec (8 hex digits), and for
// the reserved list document IDs (core.DocIDIssuers, DocIDDocs, DocIDCodecs)
// the list's own folder: `issuers`, `<Issuer>/docs` or `<Issuer>/codecs`.
func docFolder(d DocKey) string {
	switch {
	case len(d.Doc) == 2+8:
		return d.Issuer + "/" + codecsDir + "/" + d.Doc
	case d.Doc == listIDIssuers:
		return issuersDir
	case d.Doc == listIDDocs:
		return d.Issuer + "/" + docsDir
	case d.Doc == listIDCodecs:
		return d.Issuer + "/" + codecsDir
	default:
		return d.Issuer + "/" + docsDir + "/" + d.Doc
	}
}
