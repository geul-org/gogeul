//ff:func feature=core type=encoder control=sequence
//ff:what EncodeLanguageText — a UTF8-LANGUAGE-TEXT Value: the language record, then the text bytes, odd member for an odd byte count
package core

// EncodeLanguageText builds a UTF8-LANGUAGE-TEXT Value (value/text
// UTF8-LANGUAGE-TEXT): its Payload is the language Value record lang
// (LANGUAGE-SHORT-CODE or LANGUAGE-BCP47-CODE, TAG included) followed by text
// packed two bytes per WORD, and the codec is the odd member when text has an odd number
// of bytes (the last low byte 0x00), else the even one. The caller passes a valid
// language record and UTF-8 text; schema checks both (V7).
func EncodeLanguageText(lang Record, text []byte) Record {
	l, _ := EncodeRecord(lang)
	t := EncodeBytesValue(CodecUTF8LanguageTextEven, text)
	return Record{Kind: KindValue, Codec: t.Codec, Payload: append(l, t.Payload...)}
}
