//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what TestParseGeneric — raw words as they are, other tokens as records, addresses not mistaken for raw words; bad tokens fail
package notation

import (
	"testing"
)

func TestParseGeneric(t *testing.T) {
	ws, err := parseGeneric([]string{"#1", "0x0002", "0x00000000/0x0000000000000806"}, Context{})
	if err != nil || len(ws) != 2+1+9 || ws[2] != 2 {
		t.Fatalf("%04X %v", ws, err)
	}
	for _, toks := range [][]string{{"0x1"}, {"word"}} {
		if _, err := parseGeneric(toks, Context{}); err == nil {
			t.Errorf("%v: want error", toks)
		}
	}
}
