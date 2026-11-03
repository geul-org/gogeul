//ff:func feature=server type=helper control=iteration dimension=1
//ff:what labelEntries — show Issuer directories by their bound names, keeping the ID in ID
package server

import "strings"

// labelEntries renames each directory entry whose name is a labelled Issuer ID
// to that label (with the trailing /) and records the ID. URLs are unchanged.
func labelEntries(entries []indexEntry, labels map[string]string) {
	for i, e := range entries {
		id := strings.TrimSuffix(e.Name, "/")
		if n, ok := labels[id]; ok && id != e.Name {
			entries[i].Name, entries[i].ID = n+"/", id
		}
	}
}
