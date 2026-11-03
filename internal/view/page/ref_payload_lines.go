//ff:func feature=page type=helper control=iteration dimension=1
//ff:what refPayloadLines — a Reference payload one word per line, each field's first word named
package page

import "github.com/geul-org/gogeul/pkg/core"

// refPayloadLines renders the payload of Reference r with its fields named
// (refFields) instead of a single "payload" run.
func refPayloadLines(r core.Record) []HexLine {
	var out []HexLine
	pos := 0
	for _, f := range refFields(r.Codec, len(r.Payload)) {
		out = append(out, hexLines(r.Payload[pos:pos+f.N], f.Name)...)
		pos += f.N
	}
	return out
}
