//ff:func feature=core type=helper control=sequence
//ff:what IsListDocID — whether a 64-bit document ID is one of the three reserved list document IDs
package core

// IsListDocID reports whether doc is DocIDIssuers, DocIDDocs or DocIDCodecs,
// the reserved IDs of the list documents.
func IsListDocID(doc uint64) bool {
	return doc >= DocIDCodecs
}
