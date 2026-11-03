//ff:func feature=resolve type=helper control=sequence
//ff:what addrCodec — `<코덱 ID>/<tail>` under an Issuer's codecs/: an 8-digit codec ID, then the version file
package resolve

// addrCodec reads `<코덱 ID>/<name>`: the codec ID is 8 hex digits (2
// words), the name as addrTail.
func addrCodec(a Addr, id, name string, bad error) (Addr, error) {
	b, ok := parseHexID(id, 2)
	if !ok {
		return Addr{}, bad
	}
	a.Kind, a.ID = AddrCodec, hexWords(b)
	return addrTail(a, name, bad)
}
