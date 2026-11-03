//ff:func feature=core type=decoder control=sequence
//ff:what decodeLength — resolve a TAG Length code to the record's total words and header size
package core

import "errors"

// decodeLength resolves the 4-bit Length code of the TAG WORD at ws[0] to the
// record's total word count and the number of TAG words (1, 2 or 3). A
// non-canonical extension (a value the direct codes or the shorter extension
// could hold) and a truncated extension are hard errors. A 1-word extension
// below 16 words is non-canonical because the same record without the
// extension word is 14 words or fewer (direct code); a 2-word extension below
// 0x10001 words is non-canonical because the same record with the 1-word
// extension is 0xFFFF words or fewer (record.md 길이는 가장 짧은 표기).
func decodeLength(ws []Word, code uint8) (int, int, error) {
	if code < 0xE {
		return int(code) + 1, 1, nil
	}
	if code == 0xE {
		if len(ws) < 2 {
			return 0, 0, errors.New("geul/core: 1-word Length extension truncated")
		}
		total := int(ws[1])
		if total < 16 {
			return 0, 0, errors.New("geul/core: non-canonical 1-word Length extension")
		}
		return total, 2, nil
	}
	if len(ws) < 3 {
		return 0, 0, errors.New("geul/core: 2-word Length extension truncated")
	}
	total := uint64(ws[1])<<16 | uint64(ws[2])
	if total < 0x10001 {
		return 0, 0, errors.New("geul/core: non-canonical 2-word Length extension")
	}
	return int(total), 3, nil
}
