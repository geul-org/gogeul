//ff:func feature=core type=encoder control=sequence
//ff:what encodeTag — the TAG WORD plus canonical Length extension for a record with the given body size
package core

import "errors"

// maxRecordWords is the largest record a 2-word Length extension can declare.
const maxRecordWords = 0xFFFFFFFF

// encodeTag returns the TAG WORD and its Length extension for a record whose
// body (Codec extension, Codec TAG and Payload) is body words. Length is the
// record's total word count including the TAG and the extension itself, and the
// canonical form is enforced: 1-14 words use a direct code, 15-65535 the 1-word
// extension (1110), 65536-2^32-1 the 2-word extension (1111). Larger records
// are rejected.
func encodeTag(k Kind, codec uint16, body int) ([]Word, error) {
	if body < 0 {
		return nil, errors.New("geul/core: negative record body")
	}
	if 1+body <= 14 {
		return []Word{tagWord(k, uint8(body), codec)}, nil
	}
	if 2+body <= 0xFFFF {
		return []Word{tagWord(k, 0xE, codec), Word(2 + body)}, nil
	}
	total := uint64(3 + body)
	if total > maxRecordWords {
		return nil, errors.New("geul/core: record exceeds 2^32-1 words")
	}
	return []Word{tagWord(k, 0xF, codec), Word(total >> 16), Word(total)}, nil
}
