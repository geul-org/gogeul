//ff:func feature=core type=helper control=selection
//ff:what SlotMarkerName — the name of a slot marker code (C 0 ABSENT, 1 UNKNOWN-VALUE, 2 NO-VALUE)
package core

// SlotMarkerName names the slot marker of Reference C (GEUL
// reference/slot-marker): 0 ABSENT (적지 않음), 1 UNKNOWN-VALUE (알 수
// 없음), 2 NO-VALUE (없음). Any other C returns "". The number is not a
// Quantifier.
func SlotMarkerName(c uint16) string {
	switch c {
	case 0:
		return "ABSENT"
	case 1:
		return "UNKNOWN-VALUE"
	case 2:
		return "NO-VALUE"
	default:
		return ""
	}
}
