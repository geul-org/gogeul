//ff:func feature=resolve type=helper control=sequence
//ff:what addrHash — `<내용 해시>.geul` (0x + 64 lowercase hex digits) under an Issuer's hashes/
package resolve

import "strings"

// addrHash reads name, `0x<64 hex>.geul`, as a's content hash.
func addrHash(a Addr, name string, bad error) (Addr, error) {
	b, ok := parseHexID(strings.TrimSuffix(name, geulExt), 16)
	if !ok || !strings.HasSuffix(name, geulExt) {
		return Addr{}, bad
	}
	a.Kind = AddrHash
	copy(a.Hash[:], b)
	return a, nil
}
