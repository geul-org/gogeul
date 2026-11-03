//ff:type feature=core type=model
//ff:what ISSUER codec constants — the last three codes of the Reference, Value and Edge zones
package core

// The ISSUER codes are the last three codes of each Kind zone but Meta: ..DL
// carries a 2-word Issuer, ..EL a 4-word and ..FL a 16-word Issuer (the low 2
// bits of C). The extension after the TAG holds the Issuer and a 2-word codec
// number that the Issuer defines (FormatWords); the Payload follows and the
// Issuer's Codec document says what it is. Each Kind has its own Codec values
// because the zones differ in size (CodecMax).
const (
	// CodecRefIssuer2Word is REFERENCE-ISSUER-2WORD (0x05DL).
	CodecRefIssuer2Word uint16 = 0x5D
	// CodecRefIssuer4Word is REFERENCE-ISSUER-4WORD (0x05EL).
	CodecRefIssuer4Word uint16 = 0x5E
	// CodecRefIssuer16Word is REFERENCE-ISSUER-16WORD (0x05FL).
	CodecRefIssuer16Word uint16 = 0x5F
	// CodecValueIssuer2Word is VALUE-ISSUER-2WORD (0x09DL).
	CodecValueIssuer2Word uint16 = 0x3D
	// CodecValueIssuer4Word is VALUE-ISSUER-4WORD (0x09EL).
	CodecValueIssuer4Word uint16 = 0x3E
	// CodecValueIssuer16Word is VALUE-ISSUER-16WORD (0x09FL).
	CodecValueIssuer16Word uint16 = 0x3F
	// CodecEdgeIssuer2Word is EDGE-ISSUER-2WORD (0x1BDL).
	CodecEdgeIssuer2Word uint16 = 0x11D
	// CodecEdgeIssuer4Word is EDGE-ISSUER-4WORD (0x1BEL).
	CodecEdgeIssuer4Word uint16 = 0x11E
	// CodecEdgeIssuer16Word is EDGE-ISSUER-16WORD (0x1BFL).
	CodecEdgeIssuer16Word uint16 = 0x11F
	// FormatWords is the width of the codec number an Issuer defines, the same
	// for the three Kinds.
	FormatWords = 2
)
