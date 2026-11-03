//ff:func feature=resolve type=resolver control=sequence
//ff:what Dir.PublishList — publish b as the next version of a list document (issuers/, <Issuer>/docs/, <Issuer>/codecs/) under its reserved ID
package resolve

import (
	"fmt"
	"slices"

	"github.com/geul-org/gogeul/pkg/core"
)

// PublishList publishes b as the next version of the list list (one of
// core.DocIDIssuers, DocIDDocs, DocIDCodecs) of issuer: the version file in
// the list's folder, its hashes/ copy and a new versions.geul whose
// PUBLICATION is DOC-POINTER(issuer, list, version) (migration.md §1.1 95,
// 96). The Issuer list belongs to geul.org 0x00000000 only. Publishing the same
// bytes again is the version already there.
func (d Dir) PublishList(issuer []core.Word, list uint64, b []byte) (VersionEntry, error) {
	if !core.IsListDocID(list) || list == core.DocIDIssuers && !slices.Equal(issuer, []core.Word{0, 0}) {
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
