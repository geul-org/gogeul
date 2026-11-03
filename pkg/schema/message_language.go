//ff:func feature=schema type=decoder control=sequence
//ff:what messageLanguage — the language record of a MESSAGE's UTF8-LANGUAGE-TEXT, as words, when it reads
package schema

import "github.com/geul-org/gogeul/pkg/core"

// messageLanguage reads a MESSAGE Meta (grammar 0x1D9L: one UTF8-LANGUAGE-TEXT
// Value) and returns its language record (the leading LANGUAGE-SHORT-CODE or
// LANGUAGE-BCP47-CODE record) in canonical hex; false when the payload does
// not read so (its layout is the codec spec's, 2층). Two MESSAGEs of the same
// language record are the same language; a SHORT-CODE and the BCP 47 tag the
// registry gives that number are not compared here (that needs the registry).
func messageLanguage(r core.Record) (string, bool) {
	v, _, err := core.DecodeRecord(r.Payload)
	if err != nil || v.Kind != core.KindValue || v.Codec&^1 != core.CodecUTF8LanguageTextEven {
		return "", false
	}
	_, n, err := core.DecodeRecord(v.Payload)
	if err != nil {
		return "", false
	}
	return core.FormatID(v.Payload[:n]), true
}
