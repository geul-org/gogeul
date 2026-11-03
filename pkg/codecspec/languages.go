//ff:func feature=codecspec type=decoder control=iteration dimension=1
//ff:what Spec.Languages — the GEUL Language Registry: LANGUAGE-SHORT-CODE code (canonical hex) to its BCP 47 tag, from the codec spec's ENUMs
package codecspec

import "github.com/geul-org/gogeul/pkg/core"

// languageCodec is the target word of LANGUAGE-SHORT-CODE.
var languageCodec = core.TargetWord(core.KindValue, core.CodecLanguageShortCode)

// Languages returns the ENUMs of the 1-WORD raw field `code` that the
// LANGUAGE-SHORT-CODE codec spec (value/text/LANGUAGE-SHORT-CODE, codec
// 0x00000000/codecs/0x000007a0) declares: each code in canonical hex
// (core.FormatID) to the BCP 47 tag BIND names it with. It is nil when s does
// not declare that codec or its code field. The registry is a closed list, so
// a code not in it is invalid, not merely unnamed.
func (s Spec) Languages() map[string]string {
	for _, f := range s[languageCodec] {
		if f.Name == "code" && f.Width == 1 {
			return f.Enums
		}
	}
	return nil
}
