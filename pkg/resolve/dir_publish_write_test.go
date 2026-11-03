//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestDirPublishWrite — each of the three writes failing fails the publish
package resolve

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDirPublishWrite(t *testing.T) {
	b := sampleDoc(6, "x")
	ws, _ := core.ReadWords(b)
	h, _ := core.ContentHash(ws)
	k := sampleKey(6, 0).DocKey()
	for _, path := range []string{versionPath(k, 1), hashPath(geulIssuer, h), listPath(k)} {
		d := Dir{Root: t.TempDir()}
		blockWrite(t, d, path)
		if _, err := d.Publish(b); err == nil {
			t.Fatalf("%s blocked: want error", path)
		}
	}
}
