//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckBytes — even and odd codecs, empty odd payload, non-zero pad
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestCheckBytes(t *testing.T) {
	v := func(c uint16, p ...core.Word) rec {
		return recOf(core.Record{Kind: core.KindValue, Codec: c, Payload: p})
	}
	if checkBytes(recOf(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("abc")))) != nil || checkBytes(v(core.CodecUTF8TextEven)) != nil {
		t.Fatal("ok cases")
	}
	if rulesOf(checkBytes(v(core.CodecUTF8TextEven|1))) != "V4" || rulesOf(checkBytes(v(core.CodecUTF8TextEven|1, 0x6162))) != "V4" {
		t.Fatal("pad cases")
	}
	if rulesOf(checkBytes(v(core.CodecURIEven|1, 0x2000))) != "V5" {
		t.Fatal("odd URI of a space")
	}
}
