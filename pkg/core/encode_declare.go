//ff:func feature=core type=encoder control=sequence
//ff:what EncodeDeclare — a DECLARE Edge from a Declare
package core

// EncodeDeclare builds the DECLARE Edge d (GEUL t002 docs/edge/DECLARE).
func EncodeDeclare(d Declare) Record {
	p := append(localWords(d.EdgeID), d.Target)
	if len(d.Issuer) > 0 {
		p = append(append(p, d.Issuer...), Word(d.Format>>16), Word(d.Format))
	}
	return Record{Kind: KindEdge, Codec: EdgeDeclare, Payload: p}
}
