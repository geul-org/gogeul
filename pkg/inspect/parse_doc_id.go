//ff:func feature=inspect type=decoder control=sequence
//ff:what parseDocID — Meta DOC-ID as DocID: DOC-POINTER(Issuer, document, 0), or CODEC-POINTER(Issuer, codec, 0) of a codec document
package inspect

import "github.com/geul-org/gogeul/pkg/core"

// parseDocID reads Meta DOC-ID into the Issuer and the document's ID: the
// 4-word document ID of a DOC-POINTER (core.DecodeDocID), or the 2-word codec
// ID of a codec document's CODEC-POINTER (core.DecodeDocIDCodec).
func parseDocID(r core.Record) (DocID, bool) {
	if issuer, doc, err := core.DecodeDocID(r); err == nil {
		return DocID{Issuer: issuer, Doc: doc}, true
	}
	issuer, codec, err := core.DecodeDocIDCodec(r)
	if err != nil {
		return DocID{}, false
	}
	return DocID{Issuer: issuer, Doc: []core.Word{core.Word(codec >> 16), core.Word(codec)}}, true
}
