//ff:type feature=core type=model
//ff:what Meta codec constants — the 42 Meta codes (C, KKK 111) in class order
package core

// Meta codes (KKK 111, TAG 0x1C0L~0x1E9L). They run without gaps in class
// order — structure, declaration, provenance, lineage, policy, IMPORT,
// validation — and carry no class bits. Validation is last: Meta C >= 0x25
// (IsValidationMeta) is left out of the content hash, and the empty codes
// 0x1EAL~0x1FFL after it are reserved for validation. DOC-1/2/4WORD keep C
// 0/1/2, so their low 2 bits are still the width code.
const (
	MetaDoc1Word         uint16 = 0x00 // DOC-1WORD (0x1C0L)
	MetaDoc2Word         uint16 = 0x01 // DOC-2WORD (0x1C1L)
	MetaDoc4Word         uint16 = 0x02 // DOC-4WORD (0x1C2L)
	MetaBody             uint16 = 0x03 // BODY (0x1C3L)
	MetaEnd              uint16 = 0x04 // END (0x1C4L)
	MetaSpecVersion      uint16 = 0x05 // SPEC-VERSION (0x1C5L)
	MetaProfile          uint16 = 0x06 // PROFILE (0x1C6L)
	MetaDocID            uint16 = 0x07 // DOC-ID (0x1C7L)
	MetaEffectiveFrom    uint16 = 0x08 // EFFECTIVE-FROM (0x1C8L)
	MetaEffectiveUntil   uint16 = 0x09 // EFFECTIVE-UNTIL (0x1C9L)
	MetaDefaultLanguage  uint16 = 0x0A // DEFAULT-LANGUAGE (0x1CAL)
	MetaAuthor           uint16 = 0x0B // AUTHOR (0x1CBL)
	MetaCoAuthor         uint16 = 0x0C // CO-AUTHOR (0x1CCL)
	MetaCreatedAt        uint16 = 0x0D // CREATED-AT (0x1CDL)
	MetaModifiedAt       uint16 = 0x0E // MODIFIED-AT (0x1CEL)
	MetaRecordedBy       uint16 = 0x0F // RECORDED-BY (0x1CFL)
	MetaRecordedAt       uint16 = 0x10 // RECORDED-AT (0x1D0L)
	MetaParent           uint16 = 0x11 // PARENT (0x1D1L)
	MetaDerivedFrom      uint16 = 0x12 // DERIVED-FROM (0x1D2L)
	MetaGeneratedBy      uint16 = 0x13 // GENERATED-BY (0x1D3L)
	MetaTransformedBy    uint16 = 0x14 // TRANSFORMED-BY (0x1D4L)
	MetaCherryPickedFrom uint16 = 0x15 // CHERRY-PICKED-FROM (0x1D5L)
	MetaReverts          uint16 = 0x16 // REVERTS (0x1D6L)
	MetaReplaces         uint16 = 0x17 // REPLACES (0x1D7L)
	MetaSquashedFrom     uint16 = 0x18 // SQUASHED-FROM (0x1D8L)
	MetaMessage          uint16 = 0x19 // MESSAGE (0x1D9L)
	MetaTag              uint16 = 0x1A // TAG (0x1DAL)
	MetaDeltaBase        uint16 = 0x1B // DELTA-BASE (0x1DBL)
	MetaDeltaRemove      uint16 = 0x1C // DELTA-REMOVE (0x1DCL)
	MetaNextLocalID      uint16 = 0x1D // NEXT-LOCAL-ID (0x1DDL)
	MetaLicense          uint16 = 0x1E // LICENSE (0x1DEL)
	MetaAccess           uint16 = 0x1F // ACCESS (0x1DFL)
	MetaPrivacy          uint16 = 0x20 // PRIVACY (0x1E0L)
	MetaRetention        uint16 = 0x21 // RETENTION (0x1E1L)
	MetaConsent          uint16 = 0x22 // CONSENT (0x1E2L)
	MetaUsage            uint16 = 0x23 // USAGE (0x1E3L)
	MetaImportDoc        uint16 = 0x24 // IMPORT-DOC (0x1E4L)
	MetaIssuerKey        uint16 = 0x25 // ISSUER-KEY (0x1E5L), the first validation code
	MetaSignature        uint16 = 0x26 // SIGNATURE (0x1E6L)
	MetaSignedOff        uint16 = 0x27 // SIGNED-OFF (0x1E7L)
	MetaConfidence       uint16 = 0x28 // CONFIDENCE (0x1E8L)
	MetaVerdict          uint16 = 0x29 // VERDICT (0x1E9L), the last assigned Meta code
)
