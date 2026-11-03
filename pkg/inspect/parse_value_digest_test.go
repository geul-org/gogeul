//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestParseValueDigest — over every Value C the SHA-256 (16 WORD) and SHA-1 (10 WORD) digests parse only as themselves; SHORT-CODE and UUID, now an even/odd neighbour pair, are not read as text or time
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestParseValueDigest(t *testing.T) {
	for c := uint16(0); c <= 0x3F; c++ {
		v16, ok16 := ParseValue(core.Record{Kind: core.KindValue, Codec: c, Payload: make([]core.Word, 16)})
		v10, ok10 := ParseValue(core.Record{Kind: core.KindValue, Codec: c, Payload: make([]core.Word, 10)})
		if ok16 && v16.Codec != c || ok10 && v10.Codec != c {
			t.Fatalf("C 0x%02X: parsed as another codec", c)
		}
		if c == core.CodecSHA256Digest && (!ok16 || ok10) || c == core.CodecSHA1Digest && (!ok10 || ok16) {
			t.Fatalf("C 0x%02X: 16 WORD %v, 10 WORD %v", c, ok16, ok10)
		}
		if (c == core.CodecLanguageShortCode || c == core.CodecUUID || c >= 0x1E) && (ok16 || ok10) {
			t.Fatalf("C 0x%02X: no typed form, got one", c)
		}
	}
}
