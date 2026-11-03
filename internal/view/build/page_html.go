//ff:func feature=build type=renderer control=sequence
//ff:what pageHTML — the viewer page of one tree document from its bytes
package build

import (
	"io"

	"github.com/geul-org/gogeul/pkg/resolve"

	"github.com/geul-org/gogeul/internal/view/page"
	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// pageHTML writes the viewer page, with the tree browser on, of the document
// b published at URL path urlPath (without .geul) in the tree t, laid out by
// spec. It is for a static copy of the tree (build); b that is not one GEUL
// document is an error.
func pageHTML(w io.Writer, t resolve.Store, urlPath string, b []byte, spec codecspec.Spec) error {
	d, err := inspect.DecodeWith(b, spec)
	if err != nil {
		return err
	}
	p := page.Build(d, titleOf(t, urlPath))
	p.Browse = true
	return page.Render(w, p)
}
