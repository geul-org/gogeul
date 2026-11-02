//ff:func feature=cli type=helper control=selection
//ff:what packEntry — one GEULPACK item as a resolver source: an https:// set base, a local set directory or a data folder with pack sets
package cli

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/geul-org/gogeul/pkg/pack"
	"github.com/geul-org/gogeul/pkg/resolve"
)

// packEntry opens item (read-only) and names the source `pack+<item>`. Only a
// data folder has a Root: loose files beside its sets read too.
func packEntry(item string) (resolve.Entry, error) {
	ctx, name, root := context.Background(), "pack+"+item, ""
	var a *pack.Archive
	var err error
	_, setErr := os.Stat(filepath.Join(item, "latest.txt"))
	wk := filepath.Join(item, ".well-known", "geul")
	switch {
	case strings.HasPrefix(item, "https://") || strings.HasPrefix(item, "http://"):
		a, err = pack.OpenRemote(ctx, pack.HTTPStore{Base: item}, pack.Options{})
	case setErr == nil:
		a, err = pack.OpenRemote(ctx, pack.FileStore{Root: item}, pack.Options{})
	case pack.IsPacked(wk):
		a, err = pack.OpenDir(ctx, wk, packOptions())
		root = item
	default:
		err = fmt.Errorf("GEULPACK %q: no pack set or pack archive there", item)
	}
	if err != nil {
		return resolve.Entry{}, err
	}
	return resolve.Entry{Name: name, Source: resolve.Dir{Root: root, Store: a}}, nil
}
