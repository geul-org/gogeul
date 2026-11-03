//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what Split — pack an archive directory as one set per Issuer under a `.well-known/geul/` directory (`<wk>/<Issuer>/latest.txt` …)
package pack

import (
	"context"
	"os"
	"path/filepath"
	"sort"
)

// Split writes the archive directory root (resolve.Dir layout) as Issuer
// sets under wk: each Issuer's paths (its hashes/ copies included), and
// paths outside every Issuer (the Issuer list `issuers/`) with Issuer
// 0x00000000 (or the first Issuer). A set that exists gets its next generation.
func Split(ctx context.Context, root, wk string, o BuildOptions) ([]BuildResult, error) {
	paths, err := listTree(root)
	if err != nil {
		return nil, err
	}
	base := filepath.Join(root, ".well-known", "geul")
	groups := splitPaths(paths)
	names := make([]string, 0, len(groups))
	for n := range groups {
		names = append(names, n)
	}
	sort.Strings(names)
	get := func(p string) ([]byte, error) { return os.ReadFile(treeFile(base, p)) }
	var out []BuildResult
	for _, n := range names {
		res, err := Write(ctx, filepath.Join(wk, n), Batch{Put: groups[n], Full: true, Get: get}, o)
		if err != nil {
			return out, err
		}
		out = append(out, res)
	}
	return out, nil
}
