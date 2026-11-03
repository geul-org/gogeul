package core

// Reserved document IDs (4 words, the top three) of the list documents
// (migration.md §1.1 96). A list's version list (`versions.geul`) points at
// its versions by DOC-POINTER(Issuer, this ID, version ID). They are never an
// address `<Issuer>/docs/<문서 ID>/` and never in the `docs/` list, like the
// term definition index document 0.
const (
	// DocIDIssuers is the Issuer list `issuers/` (Issuer geul.org 0x00000000).
	DocIDIssuers uint64 = 0xFFFFFFFFFFFFFFFF
	// DocIDDocs is an Issuer's document list `<Issuer>/docs/`.
	DocIDDocs uint64 = 0xFFFFFFFFFFFFFFFE
	// DocIDCodecs is an Issuer's codec list `<Issuer>/codecs/`.
	DocIDCodecs uint64 = 0xFFFFFFFFFFFFFFFD
)
