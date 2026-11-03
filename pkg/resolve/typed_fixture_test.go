//ff:type feature=resolve type=model
//ff:what typedFixture — a Dir with one document (two versions), one codec, the three lists and the expected bytes of each 공표 주소
package resolve

import "github.com/geul-org/gogeul/pkg/core"

// typedFixture is what newTypedFixture publishes: d holds document
// 0xffff000000000007 (v1 "a", v2 "b"), codec 0x00000980 (v1), and version 1 of
// the Issuer, document and codec lists of 0x00000000.
type typedFixture struct {
	d                     Dir
	g, pad                []core.Word
	v1, v2, codec         []byte
	issuers, docs, codecs []byte
	list, clist           []byte
	ilist, dlist, cslist  []byte
	hash                  [32]byte
}
