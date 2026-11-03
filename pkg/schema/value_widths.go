package schema

import "github.com/geul-org/gogeul/pkg/core"

// fixedValueWidth is the payload width of fixed-width direct Value codecs
// checked by V2 (grammar Value).
var fixedValueWidth = map[uint16]int{
	core.CodecGregorianDate:     2,
	core.CodecPosixSeconds:      4,
	core.CodecDurationSeconds:   4,
	core.CodecLanguageShortCode: 1,
	core.CodecUUID:              8,
	core.CodecSHA256Digest:      16,
	core.CodecSHA1Digest:        10,
}
