//ff:type feature=core type=model
//ff:what Reference codec constants — slot markers, pointers, the 5Q schemes and CODEC-POINTER by C
package core

// Reference codes (KKK 000, C 0x00~0x3F; the Record Codec is C). C < 3 are
// the slot markers (SlotMarkerName), C 3 and 4 the ISSUER-POINTER and
// DOC-POINTER (Identified only), 5 <= C < 0x3C the eleven 5Q schemes — scheme
// C/5 with Quantifier C%5 (RefScheme, RefQuantifier) — C 0x3C the
// CODEC-POINTER and C 0x3D~0x3F the REFERENCE-ISSUER codes (codec_issuer.go).
// Schemes 1~5 identify their target directly; 6~8 (EXTERNAL-POINTER-1/2/4WORD)
// and 9~11 (EXTERNAL-ISSUER-2/4/16WORD) point at another document's LOCAL ID.
const (
	// CodecIssuerPointer is ISSUER-POINTER (0x003L): an Issuer itself; its width is the Payload.
	CodecIssuerPointer uint16 = 0x03
	// CodecDocPointer is DOC-POINTER (0x004L): Issuer, 4-word document ID and
	// 2-word version ID (0 is the latest).
	CodecDocPointer uint16 = 0x04
	// CodecCodecPointer is CODEC-POINTER (0x03CL): Issuer, 2-word codec ID and
	// 2-word version ID (0 is the latest).
	CodecCodecPointer uint16 = 0x3C
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
	CodecLocal uint16 = SchemeLocal * 5
	// CodecGeul is GEUL × Identified (0x00AL).
	CodecGeul uint16 = SchemeGeul * 5
	// CodecGeulSIDX is GEUL-SIDX × Identified (0x00FL).
	CodecGeulSIDX uint16 = SchemeGeulSIDX * 5
	// CodecWikidataQID is WIKIDATA-QID × Identified (0x014L).
	CodecWikidataQID uint16 = SchemeWikidataQID * 5
	// CodecWikidataPID is WIKIDATA-PID × Identified (0x019L).
	CodecWikidataPID uint16 = SchemeWikidataPID * 5
	// CodecExternalPointer1Word is EXTERNAL-POINTER-1WORD × Identified (0x01EL).
	CodecExternalPointer1Word uint16 = SchemeExternalPointer1Word * 5
	// CodecExternalPointer2Word is EXTERNAL-POINTER-2WORD × Identified (0x023L).
	CodecExternalPointer2Word uint16 = SchemeExternalPointer2Word * 5
	// CodecExternalPointer4Word is EXTERNAL-POINTER-4WORD × Identified (0x028L).
	CodecExternalPointer4Word uint16 = SchemeExternalPointer4Word * 5
	// CodecExternalIssuer2Word is EXTERNAL-ISSUER-2WORD × Identified (0x02DL).
	CodecExternalIssuer2Word uint16 = SchemeExternalIssuer2Word * 5
	// CodecExternalIssuer4Word is EXTERNAL-ISSUER-4WORD × Identified (0x032L).
	CodecExternalIssuer4Word uint16 = SchemeExternalIssuer4Word * 5
	// CodecExternalIssuer16Word is EXTERNAL-ISSUER-16WORD × Identified (0x037L).
	CodecExternalIssuer16Word uint16 = SchemeExternalIssuer16Word * 5
)
