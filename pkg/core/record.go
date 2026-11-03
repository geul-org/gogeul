//ff:type feature=core type=model
//ff:what Record — one GEUL record: Kind, Codec, optional ISSUER extension and Payload
package core

// Record is one decoded GEUL record. Codec is the code of the TAG WORD within
// its Kind's zone (the 9-bit code KKKKC less the zone start): 0x00~0x5F for
// Reference, 0x00~0x3F for Value and Meta, 0x00~0x11F for Edge. When Codec is an ISSUER code (the last three codes of
// the Reference, Value or Edge zone, IssuerWidth), Issuer holds the 2/4/16-word
// Issuer and Format the 2-word codec number that Issuer defines; otherwise both
// are empty. Payload is everything after the TAG WORD and its extensions. The
// Length field is not stored: it is derived from the encoded size.
type Record struct {
	Kind    Kind
	Codec   uint16
	Issuer  []Word
	Format  uint32
	Payload []Word
}
