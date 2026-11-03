//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckEmbedded — no Value, wrong codec, a Value that breaks V rules, ok
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestCheckEmbedded(t *testing.T) {
	m := func(p []core.Word) rec {
		return recOf(core.Record{Kind: core.KindMeta, Codec: core.MetaProfile, Payload: p})
	}
	if fs := checkEmbedded(m(nil), "M4", "x", isUTF8Text); rulesOf(fs) != "M4" {
		t.Fatalf("none: %+v", fs)
	}
	if fs := checkEmbedded(m(enc(core.EncodeUintValue(1))), "M4", "x", isUTF8Text); rulesOf(fs) != "M4" {
		t.Fatalf("codec: %+v", fs)
	}
	if fs := checkEmbedded(m(enc(core.Record{Kind: core.KindValue, Codec: core.CodecUTF8TextEven | 1})), "M4", "x", isUTF8Text); rulesOf(fs) != "V4" {
		t.Fatalf("value rule: %+v", fs)
	}
	if fs := checkEmbedded(m(enc(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("tangeul-1.0")))), "M4", "x", isUTF8Text); fs != nil {
		t.Fatalf("ok: %+v", fs)
	}
}
