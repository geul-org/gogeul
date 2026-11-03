//ff:func feature=schema type=rule control=selection
//ff:what checkDeltaMeta — M7: DELTA-BASE is a non-zero 2-word version, DELTA-REMOVE distinct LOCAL References, NEXT-LOCAL-ID a minimal unsigned of at least 1
package schema

import "github.com/geul-org/gogeul/pkg/core"

// checkDeltaMeta applies M7 to the diff and auto increment Meta
// (meta/lineage/DELTA-BASE, DELTA-REMOVE, NEXT-LOCAL-ID) by their core
// decoders. Where they may appear (a diff, a version list) is
// lineagePlacement's, a document-wide rule.
func checkDeltaMeta(r rec) []Fact {
	var err error
	expected := ""
	switch r.R.Codec {
	case core.MetaDeltaBase:
		_, err = core.DecodeDeltaBase(r.R)
		expected = "2 WORD 버전 ID ≥ 1"
	case core.MetaDeltaRemove:
		_, err = core.DecodeDeltaRemove(r.R)
		expected = "서로 다른 LOCAL Reference 1개 이상"
	default:
		_, err = core.DecodeNextLocalID(r.R)
		expected = "최소 폭 unsigned ≥ 1"
	}
	if err != nil {
		return []Fact{factAt("M7", Error, r, "Payload", expected, err.Error())}
	}
	return nil
}
