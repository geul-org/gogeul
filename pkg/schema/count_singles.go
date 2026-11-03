//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what countSingles — M2, M4, M5, M8: Meta that may appear at most once (MESSAGE once per language)
package schema

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// singleRule maps the Meta codecs the grammar gives "개수: 많아야 하나" to the
// rule that reports a repeat: DOC-ID M2, SPEC-VERSION·PROFILE·DEFAULT-LANGUAGE
// M4, EFFECTIVE-FROM·UNTIL M5, the others M8 (grammar 0x1C5L~0x1E9L).
var singleRule = map[uint16]string{
	core.MetaDocID: "M2", core.MetaSpecVersion: "M4", core.MetaProfile: "M4", core.MetaDefaultLanguage: "M4",
	core.MetaEffectiveFrom: "M5", core.MetaEffectiveUntil: "M5",
	core.MetaAuthor: "M8", core.MetaCreatedAt: "M8", core.MetaModifiedAt: "M8", core.MetaRecordedBy: "M8",
	core.MetaRecordedAt: "M8", core.MetaDeltaBase: "M8", core.MetaDeltaRemove: "M8", core.MetaNextLocalID: "M8",
	core.MetaLicense: "M8", core.MetaAccess: "M8", core.MetaPrivacy: "M8", core.MetaRetention: "M8",
	core.MetaConfidence: "M8", core.MetaVerdict: "M8",
}

// countSingles reports every repeat of a Meta that may appear at most once
// (singleKey); MESSAGE counts per language. HEAD Edges are not Meta and are
// never counted.
func countSingles(head []rec) []Fact {
	seen := map[string]bool{}
	var fs []Fact
	for _, r := range head {
		rule, key, single := singleKey(r)
		if single && seen[key] {
			fs = append(fs, factAt(rule, Error, r, "", "문서마다 하나", fmt.Sprintf("%s 중복", core.CodecName(core.KindMeta, r.R.Codec))))
		}
		seen[key] = true
	}
	return fs
}
