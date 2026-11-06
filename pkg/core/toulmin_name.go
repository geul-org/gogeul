//ff:func feature=core type=helper control=sequence
//ff:what ToulminName — the name of a TOULMIN direct Edge code (sections 8 and 9, Edge Codec 0x60~0x94) as the grammar index and its codec document name it
package core

// toulminNames are the names of the TOULMIN direct Edge codes (Edge Codec =
// code - 0xA0): section 8 TOULMIN 0x100L~0x119L and section 9 TOULMIN 형식
// 0x120L~0x134L (grammar README 코드 색인, codes.md).
// SUPERSEDE-CHAIN, VERDICT and EVIDENCE have names and DECLARE-only codec
// documents though their layouts are not set yet. Section 9's SOURCE, FIELD,
// ALIAS and LABEL share their names with the core codes 0x0C0L, 0x0BCL,
// 0x0A5L and 0x0A4L; names are looked up by Kind and code, never back.
var toulminNames = map[uint16]string{
	0x60: "CASE-DECLARE", 0x61: "PREDICATE-DECLARE", 0x62: "SPEC-DECLARE", 0x63: "RULE-DECLARE",
	0x64: "COUNTER-DECLARE", 0x65: "EXCEPT-DECLARE", 0x68: "ATTACKS", 0x69: "SUPERSEDES", 0x6A: "SUPERSEDE-CHAIN",
	0x70: "RUN", 0x71: "DO", 0x72: "DO-ONCE", 0x73: "UNDO", 0x74: "CHECKING", 0x75: "ENTRY", 0x76: "TRIGGER",
	0x78: "VERDICT", 0x79: "EVIDENCE",
	0x80: "SOURCE", 0x81: "TEXT", 0x82: "SECTION", 0x83: "CASE-BEGIN", 0x84: "CASE-END", 0x85: "PROSE",
	0x86: "NAMESPACE", 0x87: "FUNC", 0x88: "FIELD", 0x89: "ALIAS", 0x8A: "GROUP", 0x8B: "REQUIRES",
	0x8C: "MEANS", 0x8D: "HAS", 0x8E: "RUNS", 0x8F: "EVALUATES", 0x90: "TRIGGERS", 0x91: "UNTIL",
	0x92: "SAME-AS", 0x93: "IMPORTED-FROM", 0x94: "LABEL",
}

// ToulminName returns the name of the TOULMIN Edge code codec, or "" when
// codec is no TOULMIN code.
func ToulminName(codec uint16) string {
	return toulminNames[codec]
}
