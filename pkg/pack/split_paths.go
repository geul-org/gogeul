//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what splitPaths — assign every archive path to the set of its Issuer; paths outside every Issuer (the Issuer list) go to the home set
package pack

// splitPaths assigns every path to a set name (Issuer hex, or "" when the
// archive has no Issuer paths). `<Issuer>/…` paths, `hashes/` included, go to
// that Issuer's set; the rest (`issuers/`) to the home set (splitHome).
func splitPaths(paths []string) map[string][]string {
	seen := map[string]bool{}
	for _, p := range paths {
		if o := ownerOf(p); o != "" {
			seen[o] = true
		}
	}
	home := splitHome(seen)
	out := map[string][]string{}
	for _, p := range paths {
		owner := ownerOf(p)
		if owner == "" {
			owner = home
		}
		out[owner] = append(out[owner], p)
	}
	return out
}
