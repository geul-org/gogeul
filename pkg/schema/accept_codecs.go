package schema

import "github.com/geul-org/gogeul/pkg/core"

// Value codecs accepted inside Meta (M4~M6).
var (
	isUTF8Text = func(c uint16) bool { return c == core.CodecUTF8TextEven || c == core.CodecUTF8TextEven|1 }
	isLanguage = func(c uint16) bool {
		return c == core.CodecLanguageBCP47Even || c == core.CodecLanguageBCP47Even|1 || c == core.CodecLanguageShortCode
	}
	isTime   = func(c uint16) bool { return c == core.CodecGregorianDate || c == core.CodecPosixSeconds }
	isDigest = func(c uint16) bool { return c == core.CodecSHA256Digest }
)
