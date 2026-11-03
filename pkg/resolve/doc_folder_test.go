//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what TestDocFolder — documents under docs/, codecs under codecs/, the reserved Issuer list ID in issuers; the old docs/codecs list IDs are documents
package resolve

import "testing"

func TestDocFolder(t *testing.T) {
	for want, k := range map[string]DocKey{
		"0x00000000/docs/0x0000000000000007": {Issuer: geulIssuer, Doc: "0x0000000000000007"},
		"0x00000000/codecs/0x00000b80":       {Issuer: geulIssuer, Doc: "0x00000b80"},
		"issuers":                            {Issuer: geulIssuer, Doc: "0xffffffffffffffff"},
		"0x00000002/docs/0xfffffffffffffffe": {Issuer: "0x00000002", Doc: "0xfffffffffffffffe"},
		"0x00000002/docs/0xfffffffffffffffd": {Issuer: "0x00000002", Doc: "0xfffffffffffffffd"},
	} {
		if got := docFolder(k); got != want {
			t.Fatal(got, want)
		}
	}
}
