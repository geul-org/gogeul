//ff:func feature=schema type=rule control=sequence
//ff:what bcp47WellFormed — whether bytes are a well-formed BCP 47 language tag by the RFC 5646 ABNF (langtag, privateuse or grandfathered; case ignored)
package schema

import "regexp"

// The RFC 5646 section 2.1 ABNF as a regular expression, case-insensitive.
const (
	bcpLanguage  = `(?:[a-z]{2,3}(?:-[a-z]{3}){0,3}|[a-z]{4}|[a-z]{5,8})`
	bcpVariant   = `(?:[a-z0-9]{5,8}|[0-9][a-z0-9]{3})`
	bcpExtension = `[0-9a-wy-z](?:-[a-z0-9]{2,8})+`
	bcpPrivate   = `x(?:-[a-z0-9]{1,8})+`
	bcpLangtag   = bcpLanguage + `(?:-[a-z]{4})?(?:-(?:[a-z]{2}|[0-9]{3}))?(?:-` + bcpVariant + `)*(?:-` + bcpExtension + `)*(?:-` + bcpPrivate + `)?`
	bcpIrregular = `en-gb-oed|i-ami|i-bnn|i-default|i-enochian|i-hak|i-klingon|i-lux|i-mingo|i-navajo|i-pwn|i-tao|i-tay|i-tsu|sgn-be-fr|sgn-be-nl|sgn-ch-de`
	bcpRegular   = `art-lojban|cel-gaulish|no-bok|no-nyn|zh-guoyu|zh-hakka|zh-min|zh-min-nan|zh-xiang`
)

var bcp47RE = regexp.MustCompile(`^(?i:` + bcpLangtag + `|` + bcpPrivate + `|` + bcpIrregular + `|` + bcpRegular + `)$`)

// bcp47WellFormed reports whether b is well-formed (RFC 5646 2.2.9: it
// conforms to the ABNF), which the LANGUAGE-BCP47-CODE Value requires
// (grammar 0x076L). Validity against the IANA registry is not checked here:
// it needs a registry snapshot (2층).
func bcp47WellFormed(b []byte) bool {
	return bcp47RE.Match(b)
}
