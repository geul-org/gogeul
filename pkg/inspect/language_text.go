//ff:func feature=inspect type=decoder control=selection
//ff:what languageText — the text and language of a UTF8-LANGUAGE-TEXT: a BCP 47 tag as written, a short code as its hex
package inspect

import (
	"errors"

	"github.com/geul-org/gogeul/pkg/core"
)

// languageText decodes a UTF8-LANGUAGE-TEXT Value into its text and its
// language: a LANGUAGE-BCP47-CODE as the tag, a LANGUAGE-SHORT-CODE as its
// code in hex (the registry tag is the codec spec's, namedLanguage). Another
// language record is an error.
func languageText(r core.Record) (string, string, error) {
	lang, text, err := core.DecodeLanguageText(r)
	if err != nil {
		return "", "", err
	}
	switch {
	case lang.Codec&^1 == core.CodecLanguageBCP47Even:
		tag, err := core.DecodeBytesValue(lang, core.CodecLanguageBCP47Even)
		return string(text), string(tag), err
	case lang.Codec == core.CodecLanguageShortCode:
		return string(text), core.FormatID(lang.Payload), nil
	default:
		return "", "", errors.New("geul/inspect: UTF8-LANGUAGE-TEXT without a language record")
	}
}
