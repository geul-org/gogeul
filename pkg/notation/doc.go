// Package notation writes GEUL records as one line of text and reads them
// back, for the geul CLI (show, add, meta). A line is the record's spec name
// and its fields as tokens; an Edge line starts with its Edge Id as `#n`.
// Tokens:
//
//	ABSENT                slot marker ABSENT 0x0000, 적지 않음 (edge/CODEC 선택 필드)
//	UNKNOWN               slot marker UNKNOWN-VALUE 0x0010, 알 수 없음
//	NONE                  slot marker NO-VALUE 0x0020, 없음
//	                      (a record field takes a marker only when its CODEC
//	                      list names that exact WORD: reference/slot-marker)
//	#12                   LOCAL
//	Q42, P31              WIKIDATA-QID, WIKIDATA-PID
//	0x0000/0x…0806[@2]    DOC-POINTER (an Issuer name allowed when reading: geul.org/0x…0806)
//	0x0000/0x…0806#3      EXTERNAL-ISSUER-* (Issuer document's LOCAL)
//	"text", "text"@en     UTF8-TEXT, UTF8-LANGUAGE-TEXT (language LANGUAGE-BCP47-CODE)
//	"text"@@pt-BR         UTF8-LANGUAGE-TEXT whose language is a LANGUAGE-SHORT-CODE,
//	                      by its GEUL Language Registry tag (the 0848 codec spec's ENUM);
//	                      "text"@@0x4130 is the code itself (written so without the registry)
//	<https://…>           URI
//	1.20, -3, 12e2        SIGNED-DECIMAL, digits as written
//	NAME(arg, …)          a Value or Reference by codec name, e.g.
//	                      SIGNED-INTEGER(-3), UNSIGNED-INTEGER(5), FRACTION(2, 3), FRACTION(23.45, 1256.1)
//	                      (a FRACTION component is an INTEGER as digits, a DECIMAL with a point or exponent),
//	                      GREGORIAN-DATE(1952-03-11), QUANTITY(652, Q828224, ABSENT, ABSENT),
//	                      COORDINATE-EARTH(52.5, 13.4, 1, Q209426), COORDINATE-GEO(Q111, 4.5, 137.4, 0, ABSENT)
//	enum name             a raw field's ENUM name (TRIPLE rank: normal)
//	0x…                   raw words: a raw field, or a whole record verbatim
//
// Anything the shorter forms do not cover is written as raw words, so every
// record a line is made from reads back to the same words.
package notation
