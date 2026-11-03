//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what splitHome — the set for paths outside every Issuer: 0x0000 when present, else the first Issuer in Dir order, else the top set
package pack

import "sort"

// splitHome picks the home set among the Issuers seen.
func splitHome(owners map[string]bool) string {
	if owners["0x00000000"] {
		return "0x00000000"
	}
	ids := make([]string, 0, len(owners))
	for id := range owners {
		ids = append(ids, id)
	}
	sort.Slice(ids, func(i, j int) bool { return issuerLess(ids[i], ids[j]) })
	if len(ids) == 0 {
		return ""
	}
	return ids[0]
}
