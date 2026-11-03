//ff:func feature=toulmin type=helper control=sequence
//ff:what RecordName — the spec name of a TOULMIN direct codec
package toulmin

// recordNames maps every TOULMIN record type to its spec name (GEUL t002
// TOULMIN 코덱; 0x120L~0x134L by migration.md §3.14).
var recordNames = map[uint16]string{
	CodecCaseDeclare: "CASE-DECLARE", CodecPredicateDeclare: "PREDICATE-DECLARE",
	CodecSpecDeclare: "SPEC-DECLARE", CodecRuleDeclare: "RULE-DECLARE",
	CodecCounterDeclare: "COUNTER-DECLARE", CodecExceptDeclare: "EXCEPT-DECLARE",
	CodecAttacks: "ATTACKS", CodecSupersedes: "SUPERSEDES",
	CodecRun: "RUN", CodecDo: "DO", CodecDoOnce: "DO-ONCE",
	CodecUndo: "UNDO", CodecChecking: "CHECKING",
	CodecEntry: "ENTRY", CodecTrigger: "TRIGGER",
	CodecSource: "SOURCE", CodecText: "TEXT", CodecSection: "SECTION",
	CodecCaseBegin: "CASE-BEGIN", CodecCaseEnd: "CASE-END", CodecProse: "PROSE",
	CodecNamespace: "NAMESPACE", CodecFunc: "FUNC", CodecField: "FIELD",
	CodecAlias: "ALIAS", CodecGroup: "GROUP", CodecRequires: "REQUIRES",
	CodecMeans: "MEANS", CodecHas: "HAS", CodecRuns: "RUNS",
	CodecEvaluates: "EVALUATES", CodecTriggers: "TRIGGERS", CodecUntil: "UNTIL",
	CodecSameAs: "SAME-AS", CodecImportedFrom: "IMPORTED-FROM", CodecLabel: "LABEL",
}

// RecordName returns the spec name of ref (GEUL t002 TOULMIN 코덱), or "" when
// ref is no TOULMIN record type.
func RecordName(ref CodecRef) string {
	return recordNames[ref.Codec]
}
