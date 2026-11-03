//ff:type feature=core type=model
//ff:what ISSUER codec constants — the last three codes of the Reference, Value and Edge zones
package core

// The ISSUER codes are the last three codes of each Kind zone but Meta: ..DL
// carries a 2-word Issuer, ..EL a 4-word and ..FL a 16-word Issuer. The
// extension after the TAG holds the Issuer and a 2-word codec number that the
// Issuer defines (FormatWords); the Payload follows and the Issuer's Codec
// document says what it is. Reference and Value share the C values; Edge has
// its own because its zone is five KKK long.
const (
	// CodecRefIssuer2Word is REFERENCE-ISSUER-2WORD (0x03DL).
	CodecRefIssuer2Word uint16 = 0x3D
	// CodecRefIssuer4Word is REFERENCE-ISSUER-4WORD (0x03EL).
	CodecRefIssuer4Word uint16 = 0x3E
	// CodecRefIssuer16Word is REFERENCE-ISSUER-16WORD (0x03FL).
	CodecRefIssuer16Word uint16 = 0x3F
	// CodecValueIssuer2Word is VALUE-ISSUER-2WORD (0x07DL).
	CodecValueIssuer2Word uint16 = 0x3D
	// CodecValueIssuer4Word is VALUE-ISSUER-4WORD (0x07EL).
	CodecValueIssuer4Word uint16 = 0x3E
	// CodecValueIssuer16Word is VALUE-ISSUER-16WORD (0x07FL).
	CodecValueIssuer16Word uint16 = 0x3F
	// CodecEdgeIssuer2Word is EDGE-ISSUER-2WORD (0x1BDL).
	CodecEdgeIssuer2Word uint16 = 0x13D
	// CodecEdgeIssuer4Word is EDGE-ISSUER-4WORD (0x1BEL).
	CodecEdgeIssuer4Word uint16 = 0x13E
	// CodecEdgeIssuer16Word is EDGE-ISSUER-16WORD (0x1BFL).
	CodecEdgeIssuer16Word uint16 = 0x13F
	// FormatWords is the width of the codec number an Issuer defines, the same
	// for the three Kinds.
	FormatWords = 2
)
