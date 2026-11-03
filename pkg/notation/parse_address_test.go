//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what TestParseAddress — DOC-POINTER with and without version, EXTERNAL-ISSUER of the Issuer's width; unknown names, @ with #, and a bad LOCAL fail
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestParseAddress(t *testing.T) {
	ctx := Context{Lookup: fakeLookup}
	r, err := parseAddress("geul.org/0x0000000000000806", ctx)
	if s, _ := formatReference(r); err != nil || s != "0x00000000/0x0000000000000806" {
		t.Fatalf("%s %v", s, err)
	}
	r, _ = parseAddress("0x00000000/0x0000000000000806@2", Context{})
	if s, _ := formatReference(r); s != "0x00000000/0x0000000000000806@2" {
		t.Fatal(s)
	}
	r, err = parseAddress("geul.org/0x0000000000000806#3", ctx)
	if s, _ := formatReference(r); err != nil || s != "0x00000000/0x0000000000000806#3" {
		t.Fatalf("%s %v", s, err)
	}
	r, _ = parseAddress("self/0x0000000000000001#1", ctx)
	if r.Codec != core.CodecExternalIssuer16Word {
		t.Fatalf("%+v", r)
	}
	for _, s := range []string{"nobody/x", "0x00000000/0x0000000000000806@2#1", "geul.org/0x0000000000000806#0", "geul.org/0x0000000000000806#x", "0x00000000/0x00000980#1"} {
		if _, err := parseAddress(s, ctx); err == nil {
			t.Errorf("%s: want error", s)
		}
	}
}
