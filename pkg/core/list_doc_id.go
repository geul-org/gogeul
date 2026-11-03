package core

// Reserved document ID (4 words, the top one) of the list document
// (migration.md §1.1 96, 101). The Issuer list's version list
// (`issuers/versions.geul`) points at its versions by DOC-POINTER(geul.org
// 0x00000000, this ID, version ID). It is never an address
// `<Issuer>/docs/<문서 ID>/`, like the term definition index document 0.
const (
	// DocIDIssuers is the Issuer list `issuers/` (Issuer geul.org 0x00000000).
	DocIDIssuers uint64 = 0xFFFFFFFFFFFFFFFF
)
