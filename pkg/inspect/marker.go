//ff:type feature=inspect type=model
//ff:what Marker — a slot marker in a record field: its number (low 3 bits of Reference Codec 0000) and name
package inspect

// Marker is a slot marker filling a record field (GEUL reference/slot-marker):
// Number 0 ABSENT (적지 않음), 1 UNKNOWN-VALUE (알 수 없음), 2 NO-VALUE (없음);
// 3~7 are reserved and have no Name. It is not a Value: it points at nothing.
type Marker struct {
	Number uint16 `json:"number"`
	Name   string `json:"marker"`
}
