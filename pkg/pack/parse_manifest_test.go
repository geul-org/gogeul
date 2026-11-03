//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what TestParseManifest — header, version, sections and each required section's errors are rejected
package pack

import (
	"errors"
	"testing"
)

func TestParseManifest(t *testing.T) {
	ok := encodeManifest(sampleManifest())
	if _, err := parseManifest(ok); err != nil {
		t.Fatal(err)
	}
	mut := func(i int, v byte) []byte { b := append([]byte{}, ok...); b[i] = v; return b }
	noH := mut(64+24*5+1, 0x80) // H_LAYERS entry becomes kind 0x8006 (ignored): a required section is missing
	cases := [][]byte{ok[:20], mut(0, 'X'), mut(8, 2), mut(10, 1), mut(16, 0), noH, mut(64+24, 9)}
	for i, b := range cases {
		if _, err := parseManifest(b); !errors.Is(err, ErrCorrupt) {
			t.Fatal(i, err)
		}
	}
	m := sampleManifest()
	m.issuers = []string{"0x00000000", "0x00000000"}
	if _, err := parseManifest(encodeManifest(m)); !errors.Is(err, ErrCorrupt) {
		t.Fatal("issuers")
	}
	m = sampleManifest()
	b := encodeManifest(m)
	off := int(b[64+8]) // OBJECTS body offset (small manifest)
	b[off+4] = 9        // object kind
	if _, err := parseManifest(b); !errors.Is(err, ErrCorrupt) {
		t.Fatal("objects")
	}
	m = sampleManifest()
	m.issuers = append(m.issuers, "0x00000002")
	b = encodeManifest(m)
	b[64+24*1+16]++ // ISSUERS length one longer: trailing byte
	if _, err := parseManifest(b); !errors.Is(err, ErrCorrupt) {
		t.Fatal("issuers trailing")
	}
	for _, f := range []func(*manifest){
		func(m *manifest) { m.dicts = nil },
		func(m *manifest) { m.named = []namedEntry{{"b", Loc{}}, {"a", Loc{}}} },
		func(m *manifest) { m.pLayers = []layer{{enc: 3}} },
		func(m *manifest) { m.hLayers = []layer{{enc: 3}} },
	} {
		m := sampleManifest()
		f(m)
		b := encodeManifest(m)
		if m.dicts == nil {
			b[64+24*2+16] = 3 // DICTS length 3: not 4 + 52n
		}
		if _, err := parseManifest(b); !errors.Is(err, ErrCorrupt) {
			t.Fatal("section", err)
		}
	}
}
