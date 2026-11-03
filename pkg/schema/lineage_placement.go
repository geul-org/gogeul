//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what lineagePlacement — M7: NEXT-LOCAL-ID only in a version list, DELTA-BASE only with DOC-ID and a BODY of Edges, DELTA-REMOVE only with DELTA-BASE
package schema

import "github.com/geul-org/gogeul/pkg/core"

// lineagePlacement checks where the diff and auto increment Meta stand
// (grammar 0x1DDL NEXT-LOCAL-ID "놓는 곳: versions.geul의 HEAD", "문서 본문에는
// 두지 않는다"; 0x1DBL DELTA-BASE "DOC-ID가 있는 문서에만", "BODY의 모든 레코드가
// Edge인 문서만 diff"; 0x1DCL DELTA-REMOVE "DELTA-BASE가 있는 문서에만"). A version
// list is told by isVersionList. Each misplaced Meta is reported at its first
// occurrence.
func lineagePlacement(parts docParts) []Fact {
	has := map[uint16]rec{}
	for _, r := range parts.Head {
		if _, seen := has[r.R.Codec]; r.R.Kind == core.KindMeta && !seen {
			has[r.R.Codec] = r
		}
	}
	allEdges := true
	for _, r := range parts.Body {
		allEdges = allEdges && r.R.Kind == core.KindEdge
	}
	var fs []Fact
	if r, ok := has[core.MetaNextLocalID]; ok && !isVersionList(parts) {
		fs = append(fs, factAt("M7", Error, r, "", "버전 목록(versions.geul) HEAD에만", "버전 목록이 아닌 문서"))
	}
	base, hasBase := has[core.MetaDeltaBase]
	_, hasID := has[core.MetaDocID]
	if hasBase && !hasID {
		fs = append(fs, factAt("M7", Error, base, "", "DOC-ID가 있는 문서에만", "DOC-ID 없음"))
	}
	if hasBase && !allEdges {
		fs = append(fs, factAt("M7", Error, base, "", "BODY가 모두 Edge인 문서만 diff", "Edge 아닌 BODY 레코드"))
	}
	if r, ok := has[core.MetaDeltaRemove]; ok && !hasBase {
		fs = append(fs, factAt("M7", Error, r, "", "DELTA-BASE가 있는 문서에만", "DELTA-BASE 없음"))
	}
	return fs
}
