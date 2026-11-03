//ff:func feature=core type=decoder control=sequence
//ff:what DecodeDeclare — a DECLARE Edge as Declare: a direct target, an ISSUER code itself (target alone) or an Issuer codec (ISSUER target, Issuer, codec number)
package core

import "fmt"

// DecodeDeclare reads a DECLARE Edge (grammar 0x0BBL DECLARE): the LOCAL Edge
// Id, the target word and, for an ISSUER target followed by more words, the
// Issuer and codec number (2 words) of the Issuer codec it declares. An ISSUER
// target with nothing after it declares the ISSUER code itself, as GEUL's own
// codec document of that code does (grammar 0x0BBL 유효성 "ISSUER 코드도 그
// 코드 그대로다", migration.md §1.1 127); its Issuer is empty. Nothing else
// may follow. The target is a codec: prefix 000 and LLLL 0000 (grammar 0x0BBL
// "대상은 Codec 자리만", like a CODEC's exact target).
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
	if d.Target&0xE00F != 0 {
		return Declare{}, fmt.Errorf("geul/core: DECLARE target %04X is not a codec (prefix 000, LLLL 0000)", d.Target)
	}
	pos++
	if w := IssuerWidth(SplitTarget(d.Target)); w > 0 && pos < len(r.Payload) {
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
