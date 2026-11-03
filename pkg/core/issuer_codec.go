//ff:func feature=core type=helper control=selection
//ff:what IssuerCodec — the ISSUER Codec of a Kind for an Issuer of 2, 4 or 16 words
package core

import "fmt"

// IssuerCodec returns the ISSUER Codec of k (Reference, Value or Edge) that
// carries an Issuer of width words (2, 4 or 16). Meta and any other width are
// errors.
func IssuerCodec(k Kind, width int) (uint16, error) {
	if k == KindMeta {
		return 0, fmt.Errorf("geul/core: Meta has no ISSUER codec")
	}
	switch width {
	case 2:
		return CodecMax(k) - 2, nil
	case 4:
		return CodecMax(k) - 1, nil
	case 16:
		return CodecMax(k), nil
	default:
		return 0, fmt.Errorf("geul/core: Issuer width %d is not 2, 4 or 16 words", width)
	}
}
