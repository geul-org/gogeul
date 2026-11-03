//ff:func feature=pack type=helper control=sequence
//ff:what parseHPath — `<내용 해시>.geul` under an Issuer's hashes/ as the H key
package pack

import "strings"

// parseHPath reads name, `0x<64 hex>.geul`, into ref as an H key; any other
// name leaves ref outside the forms.
func parseHPath(ref pathRef, name string) pathRef {
	b, ok := parseHex(strings.TrimSuffix(name, ".geul"), 64)
	if !ok || !strings.HasSuffix(name, ".geul") {
		return pathRef{}
	}
	ref.space, ref.hash = spaceH, string(b)
	return ref
}
