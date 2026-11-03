//ff:func feature=core type=helper control=sequence
//ff:what IsValidationMeta — whether a Kind·Codec is a validation-class Meta (Meta C >= 0x25), left out of the content hash
package core

// IsValidationMeta reports whether a k/codec record is validation class: Kind
// Meta and C >= 0x25 (MetaIssuerKey). The validation codes are the last
// assigned Meta codes, and the empty codes after them are reserved for
// validation, so no codebook is needed to leave them out of the content hash.
func IsValidationMeta(k Kind, codec uint16) bool {
	return k == KindMeta && codec >= MetaIssuerKey
}
