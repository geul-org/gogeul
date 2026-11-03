//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what TestDocFolder — documents under docs/, codecs under codecs/, the reserved Issuer list ID in issuers; the old docs/codecs list IDs are documents
package resolve

import "testing"

func TestDocFolder(t *testing.T) {
	for want, k := range map[string]DocKey{
		"0x00000000/docs/0x0000000000000007": {Issuer: geulIssuer, Doc: "0x0000000000000007"},
		"0x00000000/codecs/0x00000B80":       {Issuer: geulIssuer, Doc: "0x00000B80"},
		"issuers":                            {Issuer: geulIssuer, Doc: "0xFFFFFFFFFFFFFFFF"},
		"0x00000002/docs/0xFFFFFFFFFFFFFFFE": {Issuer: "0x00000002", Doc: "0xFFFFFFFFFFFFFFFE"},
		"0x00000002/docs/0xFFFFFFFFFFFFFFFD": {Issuer: "0x00000002", Doc: "0xFFFFFFFFFFFFFFFD"},
	} {
		if got := docFolder(k); got != want {
			t.Fatal(got, want)
		}
	}
}
