//ff:func feature=schema type=rule control=selection
//ff:what checkMeta — M2~M7 for one HEAD Meta record; other Meta codes pass
package schema

import "github.com/geul-org/gogeul/pkg/core"

// checkMeta applies the per-record Meta rules. The reserved-code Warn M1
// was removed.
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
	default:
		return nil
	}
}
