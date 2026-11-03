package schema

// reservedListID is the document ID every Issuer reserves, in canonical hex:
// only Issuer 0x00000000 uses it, for its Issuer list `issuers/`, and only in
// the DOC-POINTERs of that list's version records (grammar 0x041L, 0x1C7L).
const reservedListID = "0xFFFFFFFFFFFFFFFF"
