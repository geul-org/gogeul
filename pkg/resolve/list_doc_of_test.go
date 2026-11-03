//ff:func feature=resolve type=encoder control=iteration dimension=1
//ff:what listDocOf — test fixture: a version list document of sample document id stating entries (deprecated and deleted marked)
package resolve

import "github.com/geul-org/gogeul/pkg/core"

// listDocOf builds the version list document of document 0xffff00000000<id>:
// one PUBLICATION per entry, then a DEPRECATION or TOMBSTONE for flagged ones.
func listDocOf(id core.Word, entries ...VersionEntry) []byte {
	doc := []core.Word{0xffff, 0, 0, id}
	var recs []core.Record
	add := func(p core.Publication) {
		p.EdgeID, p.Issuer, p.Doc = uint64(len(recs)+1), []core.Word{0, 0}, doc
		recs = append(recs, core.EncodePublication(p))
	}
	for _, e := range entries {
		add(core.Publication{Codec: core.EdgePublication, Version: e.Version, Hash: e.Hash})
	}
	for _, e := range entries {
		if e.Deprecated {
			add(core.Publication{Codec: core.EdgeDeprecation, Version: e.Version})
		}
		if e.Deleted {
			add(core.Publication{Codec: core.EdgeTombstone, Version: e.Version})
		}
	}
	return encodeVersionList(recs, nil, 0, 0)
}
