//ff:func feature=cli type=helper control=sequence
//ff:what TestCheckIssuer — a sound Issuer has no findings; a bad Issuer ID and a missing Issuer folder fail
package cli

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/resolve"
)

func TestCheckIssuer(t *testing.T) {
	h := newHarness(t)
	h.tripleSpec()
	d := resolve.Dir{Root: h.data}
	if fs, err := checkIssuer(d, "0x00000000"); err != nil || len(fs) != 0 {
		t.Fatalf("%+v %v", fs, err)
	}
	if _, err := checkIssuer(d, "bad"); err == nil {
		t.Fatal("bad ID: want error")
	}
	if _, err := checkIssuer(d, "0x00010002"); err == nil {
		t.Fatal("missing folder: want error")
	}
}
