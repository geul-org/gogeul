//ff:func feature=schema type=rule control=selection
//ff:what checkText — V5: UTF-8 validity for text (a leading U+FEFF is a character), non-empty ASCII for URI, language tags, EDTF and GEDCOM-DATE; V6 for EDTF
package schema

import (
	"unicode/utf8"

	"github.com/geul-org/gogeul/pkg/core"
)

// checkText applies V5 to the bytes of a byte-string Value whose even codec is
// even (UTF8-TEXT, URI, LANGUAGE-BCP47-CODE, EDTF, GEDCOM-DATE), then V6 to
// EDTF. The EDTF and GEDCOM date grammars are not checked yet (README 미결정).
// A text may start with U+FEFF (EF BB BF): it is a character kept as written,
// not a BOM.
func checkText(r rec, even uint16, b []byte) []Fact {
	text := even == core.CodecUTF8TextEven
	switch {
	case text && !utf8.Valid(b):
		return []Fact{factAt("V5", Error, r, "Payload", "well-formed UTF-8", "invalid UTF-8")}
	case text:
		return nil
	case len(b) == 0:
		return []Fact{factAt("V5", Error, r, "Payload", "빈 문자열이 아님", "0 byte")}
	case even == core.CodecGEDCOMDateEven && !gedcomASCII(b):
		return []Fact{factAt("V5", Error, r, "Payload", "ASCII (0x20~0x7E), 앞뒤 공백 없음", "다른 byte 또는 앞뒤 공백")}
	case even == core.CodecGEDCOMDateEven:
		return nil
	case !printableASCII(b):
		return []Fact{factAt("V5", Error, r, "Payload", "ASCII (0x21~0x7E)", "다른 byte 포함")}
	case even == core.CodecEDTFEven && edtfPlainDay(b):
		return []Fact{factAt("V6", Error, r, "Payload", "GREGORIAN-DATE로 쓰는 날짜가 아님", string(b))}
	default:
		return nil
	}
}
