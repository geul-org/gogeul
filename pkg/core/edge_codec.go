package core

// Edge codecs the core reads without a codebook (GEUL 서술 Edge,
// docs/edge/PUBLICATION, edge/BIND, edge/DECLARE): the version list records and
// the codec spec vocabulary (GEUL-ORG 0층). An Edge Codec is code - 0xA0: the
// direct codes keep their 7-bit numbers (TAG 0x0A00 + Codec<<4 + L).
const (
	// EdgeBind binds a machine-matched name to a Subject Reference (서술 0x03).
	// The Issuer list (`issuers/`) names Issuers with it; a version list holds
	// no BIND (§1.1 102).
	EdgeBind uint16 = 0x03
	// EdgeTriple is one statement: Edge Id, Predicate, Subject, Object, Rank
	// (edge/TRIPLE). Its Object may be a slot marker UNKNOWN-VALUE or NO-VALUE.
	EdgeTriple uint16 = 0x0A
	// EdgeQualifier qualifies an Edge: Edge Id, Property, Target, Value
	// (edge/QUALIFIER). Its Value may be a slot marker UNKNOWN-VALUE or NO-VALUE.
	EdgeQualifier uint16 = 0x0B
	// EdgeSource ties a record to its origin: Edge Id, Target (a BODY Edge
	// Id, or LOCAL 0 for the whole document, which only the HEAD holds),
	// Origin (a Reference, or the slot marker ABSENT or UNKNOWN-VALUE). It is
	// the only 근거 Edge (edge/SOURCE); how the record cites it goes in
	// QUALIFIERs on its Edge Id (HeadEdge).
	EdgeSource uint16 = 0x20
	// EdgePublication records one published version: Edge Id, document ID, version ID, content hash.
	EdgePublication uint16 = 0x18
	// EdgeDeprecation deprecates a published version: Edge Id, document ID, version ID.
	EdgeDeprecation uint16 = 0x19
	// EdgeTombstone records a deleted version: Edge Id, document ID, version ID, reason.
	EdgeTombstone uint16 = 0x1A
	// EdgeDeclare declares one codec's payload: Edge Id, target, FIELD list (edge/DECLARE).
	EdgeDeclare uint16 = 0x1B
	// EdgeField declares one payload field: Edge Id, DECLARE, width (edge/FIELD).
	EdgeField uint16 = 0x1C
	// EdgeEnum names one code of a raw field: Edge Id, field, code (edge/ENUM).
	EdgeEnum uint16 = 0x1D
	// EdgeCodec names one Kind·Codec a field takes: Edge Id, field, target word (edge/CODEC).
	EdgeCodec uint16 = 0x1E
)
