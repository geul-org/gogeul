//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestDirFindUnder — documents.geul, the newest, versions.geul and a version under an Issuer ID or name; bad Issuer, document name and version fail
package resolve

import (
	"bytes"
	"errors"
	"testing"
)

func TestDirFindUnder(t *testing.T) {
	d := issuersFixture(t)
	v1 := edgeDoc(7, "a")
	names, _ := d.Fetch(t.Context(), "0x0000/"+namesFile)
	list, _ := d.Fetch(t.Context(), listPath(sampleKeyOf(7)))
	const doc = "0xffff000000000007"
	for _, c := range []struct {
		issuer string
		rest   []string
		want   []byte
	}{
		{"0x0000", []string{namesFile}, names},
		{"geul.org", []string{doc + geulExt}, v1},
		{"0x0000", []string{"seven"}, v1},
		{"0x0000", []string{doc, listFile}, list},
		{"0x0000", []string{"seven", "0x00000001.geul"}, v1},
	} {
		if b, err := d.findUnder(c.issuer, c.rest); err != nil || !bytes.Equal(b, c.want) {
			t.Fatalf("%v: %v", c.rest, err)
		}
	}
	if _, err := d.findUnder("nobody", []string{doc}); !errors.Is(err, ErrNotFound) {
		t.Fatal(err)
	}
	if _, err := d.findUnder("0x0000", []string{"nothing"}); !errors.Is(err, ErrNotFound) {
		t.Fatal(err)
	}
	if _, err := d.findUnder("0x0000", []string{doc, "1"}); err == nil || errors.Is(err, ErrNotFound) {
		t.Fatal(err)
	}
}
