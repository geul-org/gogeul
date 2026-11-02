//ff:func feature=resolve type=resolver control=sequence
//ff:what localSite — the site of a `file://` folder: pack objects by pread, loose files through a plain Dir
package resolve

import "github.com/geul-org/gogeul/pkg/pack"

// localSite reads `<root>/.well-known/geul/`.
func localSite(root string) *site {
	wk := Dir{Root: root}.file("")
	return &site{objects: pack.FileStore{Root: wk}, loose: Dir{Root: root}, dir: wk, sets: map[string]*pack.Reader{}}
}
