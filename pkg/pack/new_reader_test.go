//ff:func feature=pack type=resolver control=sequence
//ff:what TestNewReader — defaults: DefaultCacheBlocks and empty caches; a given cache size is kept
package pack

import "testing"

func TestNewReader(t *testing.T) {
	r := newReader(memStore{}, Options{}, &manifest{issuers: []string{"0x00000000"}}, "n")
	if r.limit != DefaultCacheBlocks || r.inum["0x00000000"] != 0 || r.name != "n" || r.decs == nil || r.dicts == nil || r.blocks == nil {
		t.Fatal(r)
	}
	r = newReader(memStore{}, Options{CacheBlocks: 3}, &manifest{}, "n")
	if r.limit != 3 {
		t.Fatal(r)
	}
}
