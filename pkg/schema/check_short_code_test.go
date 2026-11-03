//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestCheckShortCode — a one-WORD code passes; 0x0000, 0xFFFF and a wrong width fail V2
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCheckShortCode(t *testing.T) {
	for p, want := range map[string]string{"\x11\xa7": "", "\x00\x00": "V2", "\xff\xff": "V2", "": "V2", "\x00\x01\x00\x02": "V2"} {
		ws, _ := core.ReadWords([]byte(p))
		r := recOf(core.Record{Kind: core.KindValue, Codec: core.CodecLanguageShortCode, Payload: ws})
		if got := rulesOf(checkShortCode(r)); got != want {
			t.Fatalf("%X: %q want %q", p, got, want)
		}
	}
}
