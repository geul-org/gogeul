//ff:type feature=toulmin type=model
//ff:what TOULMIN codec table — the direct Edge codes of a .geul TANGEUL document (0x0E0L~0x0FBL and, since §3.14, 0x100L~0x114L)
package toulmin

// Direct TOULMIN Edge codecs (Edge Codec = code - 0x80, TAG 0x0E0L~0x0FBL). FROZEN.
const (
	CodecCaseDeclare      uint16 = 0x60
	CodecPredicateDeclare uint16 = 0x61
	CodecSpecDeclare      uint16 = 0x62
	CodecRuleDeclare      uint16 = 0x63
	CodecCounterDeclare   uint16 = 0x64
	CodecExceptDeclare    uint16 = 0x65
	CodecAttacks          uint16 = 0x68
	CodecSupersedes       uint16 = 0x69
	CodecRun              uint16 = 0x70
	CodecDo               uint16 = 0x71
	CodecDoOnce           uint16 = 0x72
	CodecUndo             uint16 = 0x73
	CodecChecking         uint16 = 0x74
	CodecEntry            uint16 = 0x75
	CodecTrigger          uint16 = 0x76
)

// TOULMIN record types formerly written as GEUL ISSUER formats 0xC000~0xC029,
// now direct Edge codes in their old number order (migration.md §3.14; TAG
// 0x100L~0x114L, Edge Codec = code - 0x80). The family of these codes
// ((code-0x80)>>5 = 4) has no name yet.
const (
	CodecSource       uint16 = 0x80
	CodecText         uint16 = 0x81
	CodecSection      uint16 = 0x82
	CodecCaseBegin    uint16 = 0x83
	CodecCaseEnd      uint16 = 0x84
	CodecProse        uint16 = 0x85
	CodecNamespace    uint16 = 0x86
	CodecFunc         uint16 = 0x87
	CodecField        uint16 = 0x88
	CodecAlias        uint16 = 0x89
	CodecGroup        uint16 = 0x8A
	CodecRequires     uint16 = 0x8B
	CodecMeans        uint16 = 0x8C
	CodecHas          uint16 = 0x8D
	CodecRuns         uint16 = 0x8E
	CodecEvaluates    uint16 = 0x8F
	CodecTriggers     uint16 = 0x90
	CodecUntil        uint16 = 0x91
	CodecSameAs       uint16 = 0x92
	CodecImportedFrom uint16 = 0x93
	CodecLabel        uint16 = 0x94
)
