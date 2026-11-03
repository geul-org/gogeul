//ff:func feature=core type=encoder control=sequence
//ff:what prependTag — a record's TAG WORD (and Length extension) for a body of n words, followed by the body
package core

// prependTag encodes the tag of a k/codec record whose body has n words and
// appends body; n is passed apart from body so the length limit can be
// checked without such a body.
func prependTag(k Kind, codec uint16, n int, body []Word) ([]Word, error) {
	tag, err := encodeTag(k, codec, n)
	if err != nil {
		return nil, err
	}
	return append(tag, body...), nil
}
