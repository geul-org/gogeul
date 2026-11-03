//ff:func feature=resolve type=resolver control=sequence
//ff:what Dir.Overwrite — development publishing: store a document as version 1 over whatever was there, with a one-entry version list
package resolve

import (
	"errors"
	"time"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// Overwrite stores the document b as version 1 of its DOC-ID, replacing any
// earlier versions, for development while formats still change and a version
// per run would pile up (Publish keeps every version). It writes
// `<Issuer>/docs/<문서 ID>/0x00000001.geul`, `<Issuer>/hashes/<내용 해시>.geul` and a version list
// with that one PUBLICATION and no PARENT, then removes the document's other
// version files (removeOtherVersions) and the `hashes/` copies the old version
// list reached: its versions, itself and its earlier lists by PARENT
// (oldCopies, removeCopies). A version list that does not decode is replaced
// and its copies, unknown, are left.
// When version 1 already is b and the only entry, nothing is written, so a
// rerun on the same input changes no file.
func (d Dir) Overwrite(b []byte) (VersionEntry, error) {
	ws, err := core.ReadWords(b)
	if err != nil {
		return VersionEntry{}, err
	}
	h, err := core.ContentHash(ws)
	if err != nil {
		return VersionEntry{}, err
	}
	doc, _ := inspect.Decode(b)
	id, ok := docIDOf(doc)
	if !ok {
		return VersionEntry{}, errors.New("geul/resolve: publishing needs a DOC-ID")
	}
	k := Key{Issuer: id.Issuer, Doc: id.Doc}.DocKey()
	old, err := d.readList(k, k.Doc)
	next := nextLocalID(old.Next, b)
	if err == nil && len(old.Entries) == 1 && old.Entries[0].Version == 1 && old.Entries[0].Hash == h && old.Next == next {
		return old.Entries[0], d.removeOtherVersions(k)
	}
	rec := core.EncodePublication(core.Publication{Codec: core.EdgePublication, EdgeID: 1, Issuer: id.Issuer, Doc: id.Doc, Version: 1, Hash: h})
	list := encodeVersionList([]core.Record{rec}, nil, time.Now().Unix(), next)
	if err := d.writeAll(map[string][]byte{versionPath(k, 1): b, hashPath(k.Issuer, h): b, hashPath(k.Issuer, listHash(list)): list, listPath(k): list}); err != nil {
		return VersionEntry{}, err
	}
	if err := d.removeOtherVersions(k); err != nil {
		return VersionEntry{}, err
	}
	return VersionEntry{Version: 1, Hash: h}, d.removeCopies(k.Issuer, d.oldCopies(old, k), map[[32]byte]bool{h: true, listHash(list): true})
}
