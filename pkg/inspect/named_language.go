//ff:func feature=inspect type=decoder control=sequence
//ff:what namedLanguage — ParseValue, with a UTF8-LANGUAGE-TEXT's LANGUAGE-SHORT-CODE language named by the registry tag
package inspect

import "github.com/geul-org/gogeul/pkg/core"

// namedLanguage parses r (ParseValue). When r is a UTF8-LANGUAGE-TEXT whose
// language record is a LANGUAGE-SHORT-CODE and langs
// (codecspec.Spec.Languages) has that code, Lang is the registry's BCP 47 tag
// instead of the code in hex. A BCP 47 language, a code langs does not have
// and any other Value are as ParseValue gives them.
func namedLanguage(r core.Record, langs map[string]string) (Value, bool) {
	v, ok := ParseValue(r)
	if !ok || r.Codec&^1 != core.CodecUTF8LanguageTextEven || len(r.Payload) == 0 || r.Payload[0] != core.TargetWord(core.KindValue, core.CodecLanguageShortCode)|1 {
		return v, ok
	}
	if tag, named := langs[v.Lang]; named {
		v.Lang = tag
	}
	return v, ok
}
