//ff:func feature=resolve type=helper control=sequence
//ff:what issuersFixture — test fixture: an archive whose issuers.geul names Issuer 0x0000 "geul.org" and whose document 7 is published and named "seven"
package resolve

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

// issuersFixture writes issuers.geul and publishes edgeDoc(7) named "seven".
func issuersFixture(t *testing.T) Dir {
	d := Dir{Root: t.TempDir()}
	subj, _ := core.EncodeIssuerPointer([]core.Word{0, 0})
	b, _ := namesDoc(nil, nil, core.Bind{EdgeID: 1, Subject: subj, Name: []byte("geul.org")})
	d.PublishList([]core.Word{0, 0}, core.DocIDIssuers, b)
	d.Publish(edgeDoc(7, "a"))
	d.Name([]core.Word{0, 0}, []core.Word{0xffff, 0, 0, 7}, "seven")
	return d
}
