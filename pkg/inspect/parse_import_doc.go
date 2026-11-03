//ff:func feature=inspect type=decoder control=sequence
//ff:what parseImportDoc — Meta IMPORT-DOC as ImportDoc; false for a payload that is not DOC-POINTER + IMPORT ID
package inspect

import "github.com/geul-org/gogeul/pkg/core"

// parseImportDoc splits Meta IMPORT-DOC into the DOC-POINTER Reference of the
// imported document (Issuer, 4-word document ID, 2-word version ID; version 0
// is "latest"), or the CODEC-POINTER of an imported codec (2-word codec ID), and the IMPORT ID after it, the rest, a minimal unsigned integer
// (GEUL t002 docs/meta/import/IMPORT-DOC).
func parseImportDoc(r core.Record) (ImportDoc, bool) {
	ref, n, err := core.DecodeRecord(r.Payload)
	if err != nil {
		return ImportDoc{}, false
	}
	issuer, doc, ver, err := core.DecodeTargetPointer(ref)
	if err != nil || n == len(r.Payload) {
		return ImportDoc{}, false
	}
	id, err := core.DecodeUintValue(core.Record{Kind: core.KindValue, Codec: core.CodecUnsignedInteger, Payload: r.Payload[n:]})
	if err != nil {
		return ImportDoc{}, false
	}
	return ImportDoc{ID: id, Issuer: issuer, Doc: doc, Version: uint32(ver[0])<<16 | uint32(ver[1])}, true
}
