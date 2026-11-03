//ff:func feature=schema type=rule control=sequence
//ff:what checkDocument — the document-wide rules: singles, effective order, lineage Meta placement, Edge Id uniqueness over HEAD and BODY, BIND names, CODEC targets, ENUM fields and codes, DECLARE targets and FIELD order, slot markers in TRIPLE·QUALIFIER, SOURCE layout and HEAD Edges left in the BODY, NO-VALUE conflicts, version list
package schema

// checkDocument applies the rules that look at more than one record. Edge
// Ids are one numbering over the HEAD Edges and the BODY.
func checkDocument(parts docParts) []Fact {
	all := append(append([]rec{}, parts.Head...), parts.Body...)
	fs := countSingles(parts.Head)
	fs = append(fs, checkEffectiveOrder(parts.Head)...)
	fs = append(fs, lineagePlacement(parts)...)
	fs = append(fs, edgeIDDups(all)...)
	fs = append(fs, bindNames(all)...)
	fs = append(fs, codecDecls(parts.Body)...)
	fs = append(fs, enumDecls(parts.Body)...)
	fs = append(fs, declareDecls(parts.Body)...)
	fs = append(fs, statementMarkers(all)...)
	fs = append(fs, sourceEdges(parts)...)
	fs = append(fs, misplacedHeadEdges(parts)...)
	fs = append(fs, noValueConflicts(parts.Body)...)
	if isVersionList(parts) {
		fs = append(fs, checkVersionList(parts)...)
	}
	return fs
}
