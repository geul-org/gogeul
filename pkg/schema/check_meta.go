//ff:func feature=schema type=rule control=selection
//ff:what checkMeta — M1~M7 for one HEAD Meta record
package schema

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// checkMeta applies the per-record Meta rules.
func checkMeta(r rec) []Fact {
	c := r.R.Codec
	switch {
	case c == core.MetaDocID:
		return checkDocID(r)
	case c == core.MetaImportDoc:
		return checkImportDoc(r)
	case c == core.MetaSpecVersion:
		return checkSpecVersion(r)
	case c == core.MetaProfile:
		return checkEmbedded(r, "M4", "UTF8-TEXT", isUTF8Text)
	case c == core.MetaDefaultLanguage:
		return checkEmbedded(r, "M4", "LANGUAGE-BCP47-CODE 또는 LANGUAGE-SHORT-CODE", isLanguage)
	case c == core.MetaEffectiveFrom || c == core.MetaEffectiveUntil:
		return checkEmbedded(r, "M5", "GREGORIAN-DATE 또는 POSIX-TIMESTAMP-SECONDS", isTime)
	case c == core.MetaParent:
		return checkEmbedded(r, "M6", "SHA-256-DIGEST", isDigest)
	case c == core.MetaDeltaBase || c == core.MetaDeltaRemove || c == core.MetaNextLocalID:
		return checkDeltaMeta(r)
	case !metaDefined(c):
		return []Fact{factAt("M1", Warn, r, "Codec", "배정된 Meta 코드", fmt.Sprintf("0x%02X (예약)", c))}
	default:
		return nil
	}
}
