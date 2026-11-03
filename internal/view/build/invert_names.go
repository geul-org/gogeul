//ff:func feature=build type=helper control=iteration dimension=1
//ff:what invertNames — ID → name from name → ID, the first name in name order winning
package build

import "sort"

// invertNames turns name → ID into ID → name. An ID with several names shows
// the first in name order.
func invertNames(names map[string]string) map[string]string {
	keys := make([]string, 0, len(names))
	for n := range names {
		keys = append(keys, n)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	labels := map[string]string{}
	for _, n := range keys {
		labels[names[n]] = n
	}
	return labels
}
