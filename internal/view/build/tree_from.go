//ff:func feature=build type=builder control=sequence
//ff:what TreeFrom — bake an archive tree (directory or pack sets) into a static copy under dst/.well-known/geul
package build

import (
	"io/fs"
	"path/filepath"

	"github.com/geul-org/gogeul/internal/view/static"
	"github.com/geul-org/gogeul/pkg/resolve"
)

// TreeFrom writes a static copy of the archive tree t (stored at src) to
// dst/.well-known/geul, at the paths of the public address space (GEUL-ORG
// 형식 지정), for a plain web host:
//
//	<path>.geul       the document's bytes (application/vnd.geul)
//	<path>            its viewer page (text/html)
//	<dir>/index.html  the directory listing
//	<dir>/index.json  the listing for the tree browser
//	_static/          the graph scripts
//
// The codec specs come from the tree;
// a tree whose specs do not read lays records out generically.
// A document that does not decode stops the build. dst must not be src or
// inside it.
func TreeFrom(t resolve.Store, src, dst string) (Count, error) {
	b := &baker{tree: t, src: src, out: filepath.Join(dst, ".well-known", "geul")}
	if err := checkOut(b.src, b.out); err != nil {
		return Count{}, err
	}
	b.spec, _ = resolve.Dir{Store: t}.CodecSpec()
	if err := fs.WalkDir(t, ".", b.visit); err != nil {
		return b.count, err
	}
	return b.count, static.Write(filepath.Join(b.out, "_static"))
}
