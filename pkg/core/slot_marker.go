//ff:type feature=core type=model
//ff:what Slot marker WORDs — ABSENT, UNKNOWN-VALUE, NO-VALUE (Reference C 0, 1, 2 with LLLL 0000)
package core

// The slot markers (GEUL reference/slot-marker) are one-word Reference records
// of C 0, 1 and 2: TAG WORDs 0x0000, 0x0010 and 0x0020. A marker points at
// nothing; it fills a record field whose CODEC list names that exact WORD
// (CodecMatches), never a HEAD or BODY record position, an Edge Id, a field
// without CODECs or a wildcard. A marker C with LLLL >= 1 is reserved.
const (
	// Absent (적지 않음) fills an optional field left out. It says nothing
	// about the world.
	Absent Word = 0x0000
	// UnknownValue (알 수 없음) says a value exists but is not known
	// (Wikidata somevalue): ∃x, x unidentified.
	UnknownValue Word = 0x0010
	// NoValue (없음) says no value exists (Wikidata novalue): ¬∃x, closed
	// within the statement's qualifiers.
	NoValue Word = 0x0020
	// slotMarkerCodecs is the number of slot marker codes (C 0, 1, 2).
	slotMarkerCodecs uint16 = 3
)
