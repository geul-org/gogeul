//ff:func feature=resolve type=helper control=sequence
//ff:what codecDocID — the 2-word codec ID of a GEUL direct code: its target word (TAG, L=0) as 32 bits
package resolve

import "github.com/geul-org/gogeul/pkg/core"

// codecDocID returns the codec ID of the direct code in target under the GEUL
// Issuer 0x00000000 (§3.13 D24): the TAG with L=0 as 32 bits, so `0x05DL`
// is `0x000005d0`. The Kind is part of the TAG, so codec IDs do not collide.
func codecDocID(target core.Word) []core.Word {
	return []core.Word{0, target}
}
