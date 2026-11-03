//ff:func feature=schema type=rule control=sequence
//ff:what checkEdge — E1 and E4 for one Edge record
package schema

// checkEdge applies E1 (leading LOCAL Edge Id, not 0) and E4 (layout of the
// core-embedded codecs). An unassigned Edge code is not reported (the
// reserved-code Warn E3 was removed, migration.md §1.1 112).
func checkEdge(r rec) []Fact {
	var fs []Fact
	if id, ok := edgeID(r); !ok {
		fs = append(fs, factAt("E1", Error, r, "Edge Id", "Payload 첫 요소 LOCAL Reference (Quantifier 000)", "없음"))
	} else if id == 0 {
		fs = append(fs, factAt("E1", Error, r, "Edge Id", "1 이상 (0은 문서 자신)", "0"))
	}
	if want, err := builtinLayout(r.R); err != nil {
		fs = append(fs, factAt("E4", Error, r, "Payload", want, err.Error()))
	}
	return fs
}
