//ff:func feature=core type=helper control=selection
//ff:what ExternalIssuerWidth — the Issuer width of an EXTERNAL-ISSUER scheme (9, 10, 11 → 2, 4, 16)
package core

// ExternalIssuerWidth returns the Issuer width of the Reference scheme (C/5) of
// EXTERNAL-ISSUER-2/4/16WORD: scheme 9 is 2, 10 is 4 and 11 is 16 words. Any
// other scheme returns 0.
func ExternalIssuerWidth(scheme uint16) int {
	switch scheme {
	case SchemeExternalIssuer2Word:
		return 2
	case SchemeExternalIssuer4Word:
		return 4
	case SchemeExternalIssuer16Word:
		return 16
	default:
		return 0
	}
}
