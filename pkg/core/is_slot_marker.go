//ff:func feature=core type=helper control=sequence
//ff:what IsSlotMarker — whether a record is a one-word slot marker (Reference C < 3, no payload)
package core

// IsSlotMarker reports whether r is a one-word slot marker record: Reference,
// C < 3 (IsSlotMarkerCodec), no payload and no ISSUER extension. A marker C
// with a payload is not a marker; it is reserved.
func IsSlotMarker(r Record) bool {
	return r.Kind == KindReference && IsSlotMarkerCodec(r.Codec) && len(r.Payload) == 0 && len(r.Issuer) == 0
}
