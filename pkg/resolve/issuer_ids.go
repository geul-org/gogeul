//ff:func feature=resolve type=decoder control=iteration dimension=1
//ff:what issuerIDs — the Issuer IDs the Issuer list binds names to; none when it is absent or does not read
package resolve

// issuerIDs lists the IDs of the name index b, in map order.
func issuerIDs(b []byte) []string {
	names, _ := ParseIssuerNames(b)
	var out []string
	for _, id := range names {
		out = append(out, id)
	}
	return out
}
