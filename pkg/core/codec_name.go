//ff:func feature=core type=helper control=selection
//ff:what CodecName — the spec name of a Reference, Meta, Value, Edge or ISSUER codec; "" otherwise
package core

// metaNames maps the Meta codes (KKK 111) to their spec names.
var metaNames = map[uint16]string{
	MetaDoc1Word: "DOC-1WORD", MetaDoc2Word: "DOC-2WORD", MetaDoc4Word: "DOC-4WORD", MetaBody: "BODY", MetaEnd: "END",
	MetaSpecVersion: "SPEC-VERSION", MetaProfile: "PROFILE", MetaDocID: "DOC-ID", MetaEffectiveFrom: "EFFECTIVE-FROM",
	MetaEffectiveUntil: "EFFECTIVE-UNTIL", MetaDefaultLanguage: "DEFAULT-LANGUAGE",
	MetaAuthor: "AUTHOR", MetaCoAuthor: "CO-AUTHOR", MetaCreatedAt: "CREATED-AT", MetaModifiedAt: "MODIFIED-AT",
	MetaRecordedBy: "RECORDED-BY", MetaRecordedAt: "RECORDED-AT",
	MetaParent: "PARENT", MetaDerivedFrom: "DERIVED-FROM", MetaGeneratedBy: "GENERATED-BY", MetaTransformedBy: "TRANSFORMED-BY",
	MetaCherryPickedFrom: "CHERRY-PICKED-FROM", MetaReverts: "REVERTS", MetaReplaces: "REPLACES", MetaSquashedFrom: "SQUASHED-FROM",
	MetaMessage: "MESSAGE", MetaTag: "TAG", MetaDeltaBase: "DELTA-BASE", MetaDeltaRemove: "DELTA-REMOVE", MetaNextLocalID: "NEXT-LOCAL-ID",
	MetaLicense: "LICENSE", MetaAccess: "ACCESS", MetaPrivacy: "PRIVACY", MetaRetention: "RETENTION", MetaConsent: "CONSENT", MetaUsage: "USAGE",
	MetaImportDoc: "IMPORT-DOC",
	MetaIssuerKey: "ISSUER-KEY", MetaSignature: "SIGNATURE", MetaSignedOff: "SIGNED-OFF", MetaConfidence: "CONFIDENCE", MetaVerdict: "VERDICT",
}

// refNames maps the Reference C that are not quantified (markers, pointers) to
// their spec names; schemeNames names the 5Q schemes by C/5.
var refNames = map[uint16]string{
	CodecIssuerPointer: "ISSUER-POINTER", CodecDocPointer: "DOC-POINTER", CodecCodecPointer: "CODEC-POINTER",
	CodecRefIssuer2Word: "REFERENCE-ISSUER-2WORD", CodecRefIssuer4Word: "REFERENCE-ISSUER-4WORD", CodecRefIssuer16Word: "REFERENCE-ISSUER-16WORD",
}

var schemeNames = map[uint16]string{
	SchemeLocal: "LOCAL", SchemeGeul: "GEUL", SchemeGeulSIDX: "GEUL-SIDX", SchemeWikidataQID: "WIKIDATA-QID", SchemeWikidataPID: "WIKIDATA-PID",
	SchemeExternalPointer1Word: "EXTERNAL-POINTER-1WORD", SchemeExternalPointer2Word: "EXTERNAL-POINTER-2WORD",
	SchemeExternalPointer4Word: "EXTERNAL-POINTER-4WORD", SchemeExternalIssuer2Word: "EXTERNAL-ISSUER-2WORD",
	SchemeExternalIssuer4Word: "EXTERNAL-ISSUER-4WORD", SchemeExternalIssuer16Word: "EXTERNAL-ISSUER-16WORD",
}

// edgeNames maps the direct Edge codes (Codec = code - 0x80) the core knows to
// their spec names. The TOULMIN family (0x060~0x07B) and the former GEUL ISSUER
// formats now direct (0x080~0x094, TAG 0x100L~0x114L, migration.md §3.14) are
// named by pkg/toulmin.
var edgeNames = map[uint16]string{
	0x00: "REFERENCE-NODE", 0x01: "SAME-NODE", 0x02: "DIFFERENT-NODE", EdgeBind: "BIND", 0x04: "LABEL", 0x05: "ALIAS", 0x06: "DESCRIPTION",
	0x08: "VERB-SHORT", 0x09: "VERB", EdgeTriple: "TRIPLE", EdgeQualifier: "QUALIFIER", 0x10: "WORLD", 0x11: "TIMELINE", 0x12: "SCENE",
	EdgePublication: "PUBLICATION", EdgeDeprecation: "DEPRECATION", EdgeTombstone: "TOMBSTONE",
	EdgeDeclare: "DECLARE", EdgeField: "FIELD", EdgeEnum: "ENUM", EdgeCodec: "CODEC", EdgeSource: "SOURCE",
	0x41: "FUNCTION-DECLARE", 0x42: "TYPE-DECLARE",
	0x48: "CONST", 0x49: "UNARY", 0x4A: "BINARY", 0x4B: "CALL", 0x4C: "MEMBER", 0x4D: "INDEX",
	0x4E: "CONSTRUCT", 0x4F: "CONDITIONAL", 0x50: "BLOCK", 0x51: "ASSIGN", 0x52: "IF", 0x53: "SWITCH",
	0x54: "WHILE", 0x55: "FOREACH", 0x56: "JUMP", 0x57: "RETURN", 0x58: "RAISE", 0x59: "TRY",
	CodecEdgeIssuer2Word: "EDGE-ISSUER-2WORD", CodecEdgeIssuer4Word: "EDGE-ISSUER-4WORD", CodecEdgeIssuer16Word: "EDGE-ISSUER-16WORD",
}

var valueNames = map[uint16]string{
	CodecBool: "BOOL", CodecSignedInteger: "SIGNED-INTEGER", CodecUnsignedInteger: "UNSIGNED-INTEGER", CodecBinaryFloat: "BINARY-FLOAT",
	CodecCoordinateEarth: "COORDINATE-EARTH", CodecCoordinateGeo: "COORDINATE-GEO", CodecQuantity: "QUANTITY", CodecFraction: "FRACTION",
	CodecSignedDecimal: "SIGNED-DECIMAL", CodecUnsignedDecimal: "UNSIGNED-DECIMAL",
	CodecGregorianDate: "GREGORIAN-DATE", CodecCivilTime: "CIVIL-TIME-NANOSECONDS", CodecPosixSeconds: "POSIX-TIMESTAMP-SECONDS",
	CodecDurationSeconds: "DURATION-SECONDS",
	CodecEDTFEven:        "EDTF", CodecEDTFEven | 1: "EDTF", CodecGEDCOMDateEven: "GEDCOM-DATE", CodecGEDCOMDateEven | 1: "GEDCOM-DATE",
	CodecUTF8TextEven: "UTF8-TEXT", CodecUTF8TextEven | 1: "UTF8-TEXT", CodecURIEven: "URI", CodecURIEven | 1: "URI",
	CodecLanguageBCP47Even: "LANGUAGE-BCP47-CODE", CodecLanguageBCP47Even | 1: "LANGUAGE-BCP47-CODE",
	CodecUTF8LanguageTextEven: "UTF8-LANGUAGE-TEXT", CodecUTF8LanguageTextEven | 1: "UTF8-LANGUAGE-TEXT",
	CodecLanguageShortCode: "LANGUAGE-SHORT-CODE", CodecUUID: "UUID", CodecSHA256Digest: "SHA-256-DIGEST", CodecSHA1Digest: "SHA1-DIGEST",
	CodecOpaqueBytesEven: "OPAQUE-BYTES", CodecOpaqueBytesEven | 1: "OPAQUE-BYTES", CodecBitVector: "BIT-VECTOR", CodecDurationNanoseconds: "DURATION-NANOSECONDS",
	CodecValueIssuer2Word: "VALUE-ISSUER-2WORD", CodecValueIssuer4Word: "VALUE-ISSUER-4WORD", CodecValueIssuer16Word: "VALUE-ISSUER-16WORD",
}

// CodecName returns the spec name of a k/codec record: a slot marker, pointer,
// REFERENCE-ISSUER or 5Q scheme (its scheme name; QuantifierName names the
// Quantifier), a Meta, Value or Edge code. Unassigned codes return "".
func CodecName(k Kind, codec uint16) string {
	switch k {
	case KindReference:
		return refName(codec)
	case KindMeta:
		return metaNames[codec]
	case KindValue:
		return valueNames[codec]
	default:
		return edgeNames[codec]
	}
}
