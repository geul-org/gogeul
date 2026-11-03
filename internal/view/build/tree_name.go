//ff:func feature=build type=helper control=sequence
//ff:what treeName — the fs name inside the tree of a URL path under /.well-known/geul ("." for the top)
package build

import "strings"

// treeName maps a cleaned URL path to its name in the tree's fs.FS.
func treeName(p string) string {
	n := strings.Trim(strings.TrimPrefix(p, strings.TrimSuffix(wellKnown, "/")), "/")
	if n == "" {
		return "."
	}
	return n
}
