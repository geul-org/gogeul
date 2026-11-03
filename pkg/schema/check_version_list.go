//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what checkVersionList — P1~P4 over a version list document
package schema

import "fmt"

// checkVersionList applies the version list rules: no IMPORT and only
// PUBLICATION, DEPRECATION and TOMBSTONE records of the listed document (P1),
// Edge Ids 1, 2, … (P2) and the version rules (P3, P4).
func checkVersionList(parts docParts) []Fact {
	imports, _ := importsOf(parts.Head)
	var fs []Fact
	if len(imports) > 0 {
		fs = append(fs, Fact{Rule: "P1", Severity: Error, Expected: "IMPORT-DOC 없음", Actual: fmt.Sprintf("%d개", len(imports))})
	}
	st := &listState{docs: map[string]*docVersions{}}
	st.issuer, st.doc = listDoc(parts.Body)
	for i, r := range parts.Body {
		fs = append(fs, listRecord(st, i, r)...)
	}
	return fs
}
