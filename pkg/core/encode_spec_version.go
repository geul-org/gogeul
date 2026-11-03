//ff:func feature=core type=encoder control=sequence
//ff:what EncodeSpecVersion — Meta SPEC-VERSION whose payload is the version as a minimal unsigned integer
package core

// EncodeSpecVersion builds Meta SPEC-VERSION for v (GEUL t002 OVERVIEW 정의).
// The payload is the bare unsigned integer in the fewest words, with no embedded
// Value: the spec version decides how every other record is read, so it must not
// depend on a Value codec. v1 is `0x1C51 0x0001`.
func EncodeSpecVersion(v uint64) Record {
	return Record{Kind: KindMeta, Codec: MetaSpecVersion, Payload: uintWords(v)}
}
