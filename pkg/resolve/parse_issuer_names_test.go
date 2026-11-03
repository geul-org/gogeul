//ff:func feature=resolve type=decoder control=sequence
//ff:what TestParseIssuerNames — names map to Issuers, other records ignored, a later BIND wins; odd bytes and a non-document fail
package resolve

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestParseIssuerNames(t *testing.T) {
	names, err := ParseIssuerNames(issuersDoc([2]any{core.Word(0), "geul.org"}, [2]any{core.Word(7), "ietf.org"}, [2]any{core.Word(2), "geul.org"}))
	if err != nil || len(names) != 2 || names["geul.org"] != "0x00000002" || names["ietf.org"] != "0x00000007" {
		t.Fatalf("got %v %v", names, err)
	}
	if _, err := ParseIssuerNames([]byte{1}); err == nil {
		t.Fatal("odd: want error")
	}
	if _, err := ParseIssuerNames([]byte{0x1C, 0x40}); err == nil {
		t.Fatal("not a document: want error")
	}
}
