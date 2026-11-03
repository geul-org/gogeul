//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestParseValue — every known codec decodes (byte strings incl. URI, EDTF, GEDCOM-DATE, language text; coordinates; binary float; integers, FRACTION, decimal, time, digest); unknown codecs and bad payloads do not
package inspect

import (
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestParseValue(t *testing.T) {
	qty, _ := core.EncodeQuantity(core.Quantity{Amount: core.EncodeSignedDecimalBig(-2, big.NewInt(120)), Unit: core.Record{Kind: core.KindReference}, Lower: core.Record{Kind: core.KindReference}, Upper: core.Record{Kind: core.KindReference}})
	sd, _ := core.EncodeDecimalValue(true, -1, 8)
	ud, _ := core.EncodeDecimalValue(false, 0, 75)
	cases := []struct {
		r    core.Record
		want Value
	}{
		{core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("abc")), Value{Codec: core.CodecUTF8TextEven | 1, Text: "abc"}},
		{core.EncodeBytesValue(core.CodecLanguageBCP47Even, []byte("ko")), Value{Codec: core.CodecLanguageBCP47Even, Text: "ko"}},
		{core.EncodeLanguageText(core.EncodeBytesValue(core.CodecLanguageBCP47Even, []byte("ko")), []byte("우주")), Value{Codec: core.CodecUTF8LanguageTextEven, Text: "우주", Lang: "ko"}},
		{core.EncodeBytesValue(core.CodecURIEven, []byte("http://x")), Value{Codec: core.CodecURIEven, Text: "http://x"}},
		{core.EncodeBytesValue(core.CodecEDTFEven, []byte("1974")), Value{Codec: core.CodecEDTFEven, Text: "1974"}},
		{core.EncodeBytesValue(core.CodecGEDCOMDateEven, []byte("JULIAN 1300")), Value{Codec: core.CodecGEDCOMDateEven | 1, Text: "JULIAN 1300"}},
		{core.EncodeUintValue(7), Value{Codec: core.CodecUnsignedInteger, Uint: 7}},
		{core.Record{Kind: core.KindValue, Codec: core.CodecBinaryFloat, Payload: []core.Word{0x3C00}}, Value{Codec: core.CodecBinaryFloat, Float: 1}},
		{core.EncodeCoordinate(core.Coordinate{Lat: 52.5, Lon: 13.4}), Value{Codec: core.CodecCoordinateEarth, Lat: 52.5, Lon: 13.4}},
		{sd, Value{Codec: core.CodecSignedDecimal, Exp: -1, Coef: 8}},
		{ud, Value{Codec: core.CodecUnsignedDecimal, Coef: 75}},
		{core.EncodePosixSeconds(86400), Value{Codec: core.CodecPosixSeconds, Time: time.Unix(86400, 0).UTC()}},
		{core.Record{Kind: core.KindValue, Codec: core.CodecGregorianDate, Payload: []core.Word{0, 1}}, Value{Codec: core.CodecGregorianDate, Time: time.Unix(86400, 0).UTC()}},
		{core.Record{Kind: core.KindValue, Codec: core.CodecSHA256Digest, Payload: make([]core.Word, 16)}, Value{Codec: core.CodecSHA256Digest, Text: "0x" + strings.Repeat("0", 64)}},
		{core.Record{Kind: core.KindValue, Codec: core.CodecSHA1Digest, Payload: make([]core.Word, 10)}, Value{Codec: core.CodecSHA1Digest, Text: "0x" + strings.Repeat("0", 40)}},
		{qty, Value{Codec: core.CodecQuantity, Amount: "1.20"}},
		{core.Record{Kind: core.KindValue, Codec: core.CodecSignedInteger, Payload: []core.Word{0xFFFB}}, Value{Codec: core.CodecSignedInteger, Number: "-5"}},
		{core.Record{Kind: core.KindValue, Codec: core.CodecFraction, Payload: []core.Word{0x0611, 2, 0x0621, 3}}, Value{Codec: core.CodecFraction, Number: "2/3"}},
	}
	for _, c := range cases {
		if got, ok := ParseValue(c.r); !ok || got != c.want {
			t.Fatalf("got %+v %v want %+v", got, ok, c.want)
		}
	}
	for _, c := range []uint16{core.CodecUUID, core.CodecLanguageShortCode, 0x1E, 0x2A, 0x3C} {
		if _, ok := ParseValue(core.Record{Kind: core.KindValue, Codec: c}); ok {
			t.Fatalf("0x%02X: want not ok", c)
		}
	}
	if _, ok := ParseValue(core.Record{Kind: core.KindValue, Codec: core.CodecSHA256Digest, Payload: []core.Word{1}}); ok {
		t.Fatal("short digest: want not ok")
	}
	if _, ok := ParseValue(core.Record{Kind: core.KindValue, Codec: core.CodecSHA1Digest, Payload: make([]core.Word, 16)}); ok {
		t.Fatal("long SHA-1 digest: want not ok")
	}
	if _, ok := ParseValue(core.Record{Kind: core.KindValue, Codec: core.CodecUTF8TextEven | 1}); ok {
		t.Fatal("bad odd text: want not ok")
	}
	if _, ok := ParseValue(core.Record{Kind: core.KindValue, Codec: core.CodecFraction}); ok {
		t.Fatal("broken fraction: want not ok")
	}
	if _, ok := ParseValue(core.Record{Kind: core.KindValue, Codec: core.CodecBinaryFloat, Payload: make([]core.Word, 8)}); ok {
		t.Fatal("binary128: want not ok")
	}
}
