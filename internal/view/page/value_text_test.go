//ff:func feature=page type=helper control=iteration dimension=1
//ff:what TestValueText — quoted text, language text with its language, bare language tag, URI, EDTF and GEDCOM-DATE, integers, FRACTION, date, time, digest, decimal with its exact form
package page

import (
	"testing"
	"time"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestValueText(t *testing.T) {
	cases := map[string]inspect.Value{
		`"a\n"`:                {Codec: core.CodecUTF8TextEven | 1, Text: "a\n"},
		"ko":                   {Codec: core.CodecLanguageBCP47Even, Text: "ko"},
		`"우주"@ko`:              {Codec: core.CodecUTF8LanguageTextEven, Text: "우주", Lang: "ko"},
		"http://x":             {Codec: core.CodecURIEven | 1, Text: "http://x"},
		"197X":                 {Codec: core.CodecEDTFEven, Text: "197X"},
		"JULIAN 1300":          {Codec: core.CodecGEDCOMDateEven | 1, Text: "JULIAN 1300"},
		"7":                    {Codec: core.CodecUnsignedInteger, Uint: 7},
		"4.5":                  {Codec: core.CodecBinaryFloat, Float: 4.5},
		"1, 2 ±0.1 Q11573":     {Codec: core.CodecCoordinateEarth, Lat: 1, Lon: 2, Float: float64(float32(0.1)), Unit: "Q11573"},
		"0.8 (8e-1)":           {Codec: core.CodecSignedDecimal, Coef: 8, Exp: -1},
		"750 (75e1)":           {Codec: core.CodecUnsignedDecimal, Coef: 75, Exp: 1},
		"1970-01-02":           {Codec: core.CodecGregorianDate, Time: time.Unix(86400, 0).UTC()},
		"1970-01-02T00:00:00Z": {Codec: core.CodecPosixSeconds, Time: time.Unix(86400, 0).UTC()},
		"0x00ff":               {Codec: core.CodecSHA256Digest, Text: "0x00ff"},
		"652 Q1 [651, 653]":    {Codec: core.CodecQuantity, Amount: "652", Unit: "Q1", Lower: "651", Upper: "653"},
		"0x01ff":               {Codec: core.CodecSHA1Digest, Text: "0x01ff"},
		"-5":                   {Codec: core.CodecSignedInteger, Number: "-5"},
		"23.45/1256.1":         {Codec: core.CodecFraction, Number: "23.45/1256.1"},
	}
	for want, v := range cases {
		if got := valueText(v); got != want {
			t.Fatalf("got %q want %q", got, want)
		}
	}
}
