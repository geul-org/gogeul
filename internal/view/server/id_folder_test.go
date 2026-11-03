//ff:func feature=server type=helper control=sequence
//ff:what TestIDFolder — an 8-digit codec ID lives under codecs, a 16-digit document ID under docs
package server

import "testing"

func TestIDFolder(t *testing.T) {
	if idFolder("0x00000980") != "codecs" || idFolder("0x0000000000000007") != "docs" {
		t.Fatal()
	}
}
