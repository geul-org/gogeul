//ff:func feature=toulmin type=helper control=selection
//ff:what RecordName — the spec name of a TOULMIN direct codec or GEUL format number
package toulmin

// recordNames maps every TOULMIN record type to its spec name (GEUL t002
// TOULMIN 코덱). Direct codecs are keyed by codec, formats by format number.
var recordNames = map[CodecRef]string{
	{Codec: CodecCaseDeclare}: "CASE-DECLARE", {Codec: CodecPredicateDeclare}: "PREDICATE-DECLARE",
	{Codec: CodecSpecDeclare}: "SPEC-DECLARE", {Codec: CodecRuleDeclare}: "RULE-DECLARE",
	{Codec: CodecCounterDeclare}: "COUNTER-DECLARE", {Codec: CodecExceptDeclare}: "EXCEPT-DECLARE",
	{Codec: CodecAttacks}: "ATTACKS", {Codec: CodecSupersedes}: "SUPERSEDES",
	{Codec: CodecRun}: "RUN", {Codec: CodecDo}: "DO", {Codec: CodecDoOnce}: "DO-ONCE",
	{Codec: CodecUndo}: "UNDO", {Codec: CodecChecking}: "CHECKING",
	{Codec: CodecEntry}: "ENTRY", {Codec: CodecTrigger}: "TRIGGER",
}

// formatNames maps the GEUL ISSUER format numbers to their spec names.
var formatNames = map[uint32]string{
	FormatSource: "SOURCE", FormatText: "TEXT", FormatSection: "SECTION",
	FormatCaseBegin: "CASE-BEGIN", FormatCaseEnd: "CASE-END", FormatProse: "PROSE",
	FormatNamespace: "NAMESPACE", FormatFunc: "FUNC", FormatField: "FIELD",
	FormatAlias: "ALIAS", FormatGroup: "GROUP", FormatRequires: "REQUIRES",
	FormatMeans: "MEANS", FormatHas: "HAS", FormatRuns: "RUNS",
	FormatEvaluates: "EVALUATES", FormatTriggers: "TRIGGERS", FormatUntil: "UNTIL",
	FormatSameAs: "SAME-AS", FormatImportedFrom: "IMPORTED-FROM", FormatLabel: "LABEL",
}

// RecordName returns the spec name of ref (GEUL t002 TOULMIN 코덱), or "" when
// ref is no TOULMIN record type.
func RecordName(ref CodecRef) string {
	switch ref.Format {
	case 0:
		return recordNames[ref]
	default:
		return formatNames[ref.Format]
	}
}
