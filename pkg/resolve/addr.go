//ff:type feature=resolve type=model
//ff:what Addr — one 공표 주소: the base it is read from, the form, the Issuer, the codec or document ID, the version or the hash
package resolve

import "github.com/geul-org/gogeul/pkg/core"

// Addr is one of the ten 공표 주소 forms (§3.13), parsed (ParseAddr).
// Base is where it is read from (`file://<dir>`, `https://<host>`), "" when the
// address was given relative to `.well-known/geul/`. Issuer is canonical (a
// zero-padded 16- or 64-digit form names the same Issuer). ID is the 2-word
// codec ID of AddrCodec or the 4-word document ID of AddrDoc. Versions asks
// for the version list (`versions.geul`); otherwise Version is the version ID,
// Latest (0) the newest (`latest.geul`). Hash is the content hash of AddrHash.
type Addr struct {
	Base     string
	Kind     AddrKind
	Issuer   []core.Word
	ID       []core.Word
	Version  uint32
	Versions bool
	Hash     [32]byte
}
