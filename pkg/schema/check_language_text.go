//ff:func feature=schema type=rule control=sequence
//ff:what checkLanguageText — V7: a UTF8-LANGUAGE-TEXT starts with a whole language record that is itself valid, then UTF-8 text with odd padding
package schema

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// checkLanguageText applies V7 to a UTF8-LANGUAGE-TEXT Value: its Payload
// starts with one whole LANGUAGE-SHORT-CODE or LANGUAGE-BCP47-CODE record,
// which must pass its own Value rules (checkValue) and is not the BCP 47
// tag zxx, no linguistic content (grammar 0x078L: such a string is a
// UTF8-TEXT; a SHORT-CODE's tag needs the registry, 2층), and the words after it
// are the text, which follows V4 (odd padding) and V5 (UTF-8, a leading U+FEFF allowed) as
// UTF8-TEXT does.
func checkLanguageText(r rec) []Fact {
	lang, n, err := core.DecodeRecord(r.R.Payload)
	if err != nil || lang.Kind != core.KindValue || !isLanguage(lang.Codec) {
		return []Fact{factAt("V7", Error, r, "Payload", "LANGUAGE-SHORT-CODE 또는 LANGUAGE-BCP47-CODE 레코드로 시작", fmt.Sprintf("%v", err))}
	}
	if fs := checkValue(rec{Index: r.Index, Word: r.Word, Total: n, R: lang}); len(fs) > 0 {
		return fs
	}
	if lang.Codec != core.CodecLanguageShortCode && noLinguisticContent(core.PutWords(lang.Payload), lang.Codec&1 == 1) {
		return []Fact{factAt("V7", Error, r, "Payload", "zxx(언어 내용 없음)가 아닌 언어 (그런 문자열은 UTF8-TEXT)", "zxx")}
	}
	text := r
	text.R = core.Record{Kind: core.KindValue, Codec: core.CodecUTF8TextEven | r.R.Codec&1, Payload: r.R.Payload[n:]}
	return checkBytes(text)
}
