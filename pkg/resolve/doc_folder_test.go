//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what TestDocFolder — documents under docs/, codecs under codecs/, the reserved list IDs in their list folders
package resolve

import "testing"

func TestDocFolder(t *testing.T) {
	for want, k := range map[string]DocKey{
		"0x00000000/docs/0x0000000000000007": {Issuer: geulIssuer, Doc: "0x0000000000000007"},
		"0x00000000/codecs/0x00000980":       {Issuer: geulIssuer, Doc: "0x00000980"},
		"issuers":                            {Issuer: geulIssuer, Doc: "0xffffffffffffffff"},
		"0x00000002/docs":                    {Issuer: "0x00000002", Doc: "0xfffffffffffffffe"},
		"0x00000002/codecs":                  {Issuer: "0x00000002", Doc: "0xfffffffffffffffd"},
	} {
		if got := docFolder(k); got != want {
			t.Fatal(got, want)
		}
	}
}
