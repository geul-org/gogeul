package core

// Reserved document IDs (4 words) of every Issuer: the
// bottom and the top one.
//
// DocIDIssuers is used only by geul.org (0x00000000) for its Issuer list
// `issuers/` (§1.1 96, 101); its version list (`issuers/versions.geul`) points
// at its versions by DOC-POINTER(0x00000000, this ID, version ID). Under any
// Issuer `<Issuer>/docs/0xFFFFFFFFFFFFFFFF/` is no address (404) and a
// document with this ID is not published as an ordinary document.
//
// DocIDTermIndex is the place of an Issuer's term definition index
// (recommended). It is published and read like any other document.
const (
	DocIDIssuers   uint64 = 0xFFFFFFFFFFFFFFFF
	DocIDTermIndex uint64 = 0x0000000000000000
)
