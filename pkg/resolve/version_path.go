//ff:func feature=resolve type=helper control=sequence
//ff:what versionPath — `<Issuer>/docs/<문서 ID>/<버전 ID>.geul` (codecs/<코덱 ID>/ for a codec), a version's address
package resolve

// versionPath is the immutable canonical address of version v of d, relative
// to `.well-known/geul/` (§3.13): `<Issuer>/docs/<문서 ID>/<버전 ID>.geul`, or
// `<Issuer>/codecs/<코덱 ID>/<버전 ID>.geul` when d is a codec (docFolder).
func versionPath(d DocKey, v uint32) string {
	return docFolder(d) + "/" + FormatVersion(v) + geulExt
}
