//ff:func feature=resolve type=resolver control=sequence
//ff:what Dir.PublishList — publish b as the next version of the list document issuers/ under its reserved ID (geul.org only)
package resolve

import (
	"fmt"
	"slices"

	"github.com/geul-org/gogeul/pkg/core"
)

// PublishList publishes b as the next version of the list list of issuer: the
// version file in the list's folder, its hashes/ copy and a new versions.geul
// whose PUBLICATION is DOC-POINTER(issuer, list, version) (migration.md §1.1
// 95, 96). The one list is the Issuer list (core.DocIDIssuers) of geul.org
// 0x00000000; there is no document or codec list (§1.1 101). Publishing the
// same bytes again is the version already there.
func (d Dir) PublishList(issuer []core.Word, list uint64, b []byte) (VersionEntry, error) {
	if list != core.DocIDIssuers || !slices.Equal(issuer, []core.Word{0, 0}) {
		return VersionEntry{}, fmt.Errorf("geul/resolve: 0x%016x of %s is not a list document", list, core.FormatID(issuer))
	}
	ws, err := core.ReadWords(b)
	if err != nil {
		return VersionEntry{}, err
	}
	h, err := core.ContentHash(ws)
	if err != nil {
		return VersionEntry{}, err
	}
	doc := uintWords4(list)
	return d.publishAt(Key{Issuer: issuer, Doc: doc}.DocKey(), issuer, doc, b, h)
}
