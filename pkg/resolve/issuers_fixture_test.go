//ff:func feature=resolve type=helper control=sequence
//ff:what issuersFixture — test fixture: an archive whose Issuer list issuers/ names Issuer 0x00000000 "geul.org" and whose document 7 is published
package resolve

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

// issuersFixture publishes the Issuer list and edgeDoc(7).
func issuersFixture(t *testing.T) Dir {
	d := Dir{Root: t.TempDir()}
	subj, _ := core.EncodeIssuerPointer([]core.Word{0, 0})
	d.PublishList([]core.Word{0, 0}, core.DocIDIssuers, bindDoc(core.Bind{Subject: subj, Name: []byte("geul.org")}))
	d.Publish(edgeDoc(7, "a"))
	return d
}
