//ff:func feature=core type=decoder control=sequence
//ff:what DecodeDeclare — a DECLARE Edge as Declare
package core

import "fmt"

// DecodeDeclare reads a DECLARE Edge (GEUL t002 docs/edge/DECLARE): the LOCAL
// Edge Id, the target word and, for an ISSUER target, the Issuer and format
// number (2 words). Nothing may follow.
func DecodeDeclare(r Record) (Declare, error) {
	if r.Kind != KindEdge || r.Codec != EdgeDeclare {
		return Declare{}, fmt.Errorf("geul/core: not a DECLARE Edge (kind %d codec 0x%02X)", r.Kind, r.Codec)
	}
	id, pos, err := takeLocal(r.Payload)
	if err != nil {
		return Declare{}, err
	}
	if pos >= len(r.Payload) {
		return Declare{}, fmt.Errorf("geul/core: DECLARE without a target")
	}
	d := Declare{EdgeID: id, Target: r.Payload[pos]}
	pos++
	if w := IssuerWidth(SplitTarget(d.Target)); w > 0 {
		if len(r.Payload) < pos+w+FormatWords {
			return Declare{}, fmt.Errorf("geul/core: DECLARE ISSUER target needs %d Issuer words and a 2-word codec number", w)
		}
		d.Issuer = r.Payload[pos : pos+w]
		d.Format = uint32(r.Payload[pos+w])<<16 | uint32(r.Payload[pos+w+1])
		pos += w + FormatWords
	}
	if pos != len(r.Payload) {
		return Declare{}, fmt.Errorf("geul/core: DECLARE has %d words after its target", len(r.Payload)-pos)
	}
	return d, nil
}
