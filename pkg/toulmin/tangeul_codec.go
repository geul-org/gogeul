//ff:type feature=toulmin type=model
//ff:what TOULMIN codec table — the direct Edge codes and GEUL ISSUER format numbers of a .geul TANGEUL document
package toulmin

// IssuerGEUL is the 2-word GEUL Issuer (0x00000000) under which the TANGEUL
// ISSUER-format records are defined (EDGE-ISSUER-2WORD, grammar TOULMIN 코덱).
const IssuerGEUL uint32 = 0x00000000

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

// TOULMIN codec numbers under the GEUL Issuer (EDGE-ISSUER-2WORD 0x1BDL, 2-word
// codec number). The values are FROZEN append-only; only their width grew.
const (
	FormatSource       uint32 = 0x0000C000
	FormatText         uint32 = 0x0000C001
	FormatSection      uint32 = 0x0000C002
	FormatCaseBegin    uint32 = 0x0000C003
	FormatCaseEnd      uint32 = 0x0000C004
	FormatProse        uint32 = 0x0000C005
	FormatNamespace    uint32 = 0x0000C010
	FormatFunc         uint32 = 0x0000C011
	FormatField        uint32 = 0x0000C012
	FormatAlias        uint32 = 0x0000C013
	FormatGroup        uint32 = 0x0000C014
	FormatRequires     uint32 = 0x0000C020
	FormatMeans        uint32 = 0x0000C021
	FormatHas          uint32 = 0x0000C022
	FormatRuns         uint32 = 0x0000C023
	FormatEvaluates    uint32 = 0x0000C024
	FormatTriggers     uint32 = 0x0000C025
	FormatUntil        uint32 = 0x0000C026
	FormatSameAs       uint32 = 0x0000C027
	FormatImportedFrom uint32 = 0x0000C028
	FormatLabel        uint32 = 0x0000C029
)
