//ff:func feature=resolve type=helper control=sequence
//ff:what codecKey — `0x00000000/codecs/<codec ID>`, the address of a direct codec's codec document
package resolve

import "github.com/geul-org/gogeul/pkg/core"

// codecKey is the address of the codec document of the codec in target
// (GEUL-ORG 직접 코드 코덱북): a codec of the GEUL Issuer 0x00000000 whose
// codec ID is the target word as 32 bits (codecDocID).
func codecKey(target core.Word) DocKey {
	return DocKey{Issuer: geulIssuer, Doc: core.FormatID(codecDocID(target))}
}
