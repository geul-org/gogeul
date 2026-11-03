//ff:type feature=core type=model
//ff:what Reference codec constants — slot markers, pointers, the 5Q schemes and CODEC-POINTER by C (three sections)
package core

// Reference codes (sections 0~2, TAG 0x000L~0x05FL; the Record Codec is the
// 9-bit code, C 0x00~0x5F). Section 0 holds the entities: the slot markers
// C < 3 (SlotMarkerName), the empty 0x03·0x04, the five entity schemes 1~5
// × 5Q at C 0x05~0x1D and the empty 0x1E·0x1F. Section 1 holds EXTERNAL: the
// schemes 6~11 × 5Q at C 0x20~0x3D, empty 0x3E·0x3F. Section 2 holds
// ISSUER-POINTER 0x40, DOC-POINTER 0x41 and CODEC-POINTER 0x42 (Identified
// only), the empty 0x43~0x5C and REFERENCE-ISSUER 0x5D~0x5F
// (codec_issuer.go). A 5Q code is read through its section KKKK and offset C
// (refSlot, RefScheme, RefQuantifier). Schemes 1~5 identify their target
// directly; 6~8 (EXTERNAL-POINTER-1/2/4WORD) and 9~11
// (EXTERNAL-ISSUER-2/4/16WORD) point at another document's LOCAL ID.
const (
	// CodecIssuerPointer is ISSUER-POINTER (0x040L): an Issuer itself; its width is the Payload.
	CodecIssuerPointer uint16 = 0x40
	// CodecDocPointer is DOC-POINTER (0x041L): Issuer, 4-word document ID and
	// 2-word version ID (0 is the latest).
	CodecDocPointer uint16 = 0x41
	// CodecCodecPointer is CODEC-POINTER (0x042L): Issuer, 2-word codec ID and
	// 2-word version ID (0 is the latest).
	CodecCodecPointer uint16 = 0x42
	// SchemeLocal is an ID this document issues (Edge Ids included).
	SchemeLocal uint16 = 1
	// SchemeGeul is an entity geul.org issues under CC0.
	SchemeGeul uint16 = 2
	// SchemeGeulSIDX is a GEUL Semantic-aligned Index (SIDX).
	SchemeGeulSIDX uint16 = 3
	// SchemeWikidataQID is a Wikidata item (the integer after Q).
	SchemeWikidataQID uint16 = 4
	// SchemeWikidataPID is a Wikidata property (the integer after P).
	SchemeWikidataPID uint16 = 5
	// SchemeExternalPointer1Word points at an imported document's LOCAL ID (1-word IMPORT ID).
	SchemeExternalPointer1Word uint16 = 6
	// SchemeExternalPointer2Word points with a 2-word IMPORT ID.
	SchemeExternalPointer2Word uint16 = 7
	// SchemeExternalPointer4Word points with a 4-word IMPORT ID.
	SchemeExternalPointer4Word uint16 = 8
	// SchemeExternalIssuer2Word cites a 2-word Issuer document's LOCAL ID.
	SchemeExternalIssuer2Word uint16 = 9
	// SchemeExternalIssuer4Word cites a 4-word Issuer document's LOCAL ID.
	SchemeExternalIssuer4Word uint16 = 10
	// SchemeExternalIssuer16Word cites a 16-word Issuer document's LOCAL ID.
	SchemeExternalIssuer16Word uint16 = 11
	// CodecLocal is LOCAL × Identified (0x005L): the form every Edge Id
	// declaration uses.
	CodecLocal uint16 = 0x05
	// CodecGeul is GEUL × Identified (0x00AL).
	CodecGeul uint16 = 0x0A
	// CodecGeulSIDX is GEUL-SIDX × Identified (0x00FL).
	CodecGeulSIDX uint16 = 0x0F
	// CodecWikidataQID is WIKIDATA-QID × Identified (0x014L).
	CodecWikidataQID uint16 = 0x14
	// CodecWikidataPID is WIKIDATA-PID × Identified (0x019L).
	CodecWikidataPID uint16 = 0x19
	// CodecExternalPointer1Word is EXTERNAL-POINTER-1WORD × Identified (0x020L).
	CodecExternalPointer1Word uint16 = 0x20
	// CodecExternalPointer2Word is EXTERNAL-POINTER-2WORD × Identified (0x025L).
	CodecExternalPointer2Word uint16 = 0x25
	// CodecExternalPointer4Word is EXTERNAL-POINTER-4WORD × Identified (0x02AL).
	CodecExternalPointer4Word uint16 = 0x2A
	// CodecExternalIssuer2Word is EXTERNAL-ISSUER-2WORD × Identified (0x02FL).
	CodecExternalIssuer2Word uint16 = 0x2F
	// CodecExternalIssuer4Word is EXTERNAL-ISSUER-4WORD × Identified (0x034L).
	CodecExternalIssuer4Word uint16 = 0x34
	// CodecExternalIssuer16Word is EXTERNAL-ISSUER-16WORD × Identified (0x039L).
	CodecExternalIssuer16Word uint16 = 0x39
)
