//ff:func feature=schema type=rule control=sequence
//ff:what checkShortCode — V2: a LANGUAGE-SHORT-CODE is one WORD and neither 0x0000 nor 0xFFFF
package schema

import "fmt"

// checkShortCode applies V2 to a LANGUAGE-SHORT-CODE: one WORD (checkWidth),
// and never 0x0000 or 0xFFFF, which the registry never assigns and which are
// invalid in this Value (grammar 0x07AL). Whether the registry assigns the
// value is the codec spec's (2층).
func checkShortCode(r rec) []Fact {
	if fs := checkWidth(r, 1); fs != nil {
		return fs
	}
	if p := r.R.Payload[0]; p == 0 || p == 0xFFFF {
		return []Fact{factAt("V2", Error, r, "Payload", "0x0000·0xFFFF가 아닌 코드", fmt.Sprintf("0x%04X", uint16(p)))}
	}
	return nil
}
