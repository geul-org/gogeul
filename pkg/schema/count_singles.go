//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what countSingles — M2, M4, M5: Meta that may appear at most once
package schema

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// singleRule maps Meta codecs that may appear once to the rule they belong to.
var singleRule = map[uint16]string{
	core.MetaDocID: "M2", core.MetaSpecVersion: "M4", core.MetaProfile: "M4", core.MetaDefaultLanguage: "M4",
	core.MetaEffectiveFrom: "M5", core.MetaEffectiveUntil: "M5",
}

// countSingles reports every repeat of a Meta that may appear at most once.
func countSingles(head []rec) []Fact {
	seen := map[uint16]bool{}
	var fs []Fact
	for _, r := range head {
		rule, single := singleRule[r.R.Codec]
		if single && seen[r.R.Codec] {
			fs = append(fs, factAt(rule, Error, r, "", "문서마다 하나", fmt.Sprintf("%s 중복", core.CodecName(core.KindMeta, r.R.Codec))))
		}
		seen[r.R.Codec] = true
	}
	return fs
}
