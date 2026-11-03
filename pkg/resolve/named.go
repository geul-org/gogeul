//ff:type feature=resolve type=model
//ff:what Named — an Issuer or document ID found in an archive with the names bound to it (Issuers only), main name last
package resolve

// Named is an ID in canonical hex and its names in binding order; the last is
// the main name. Only Issuers have names (the Issuer list, GEUL-ORG Issuer
// 이름); documents have none (§1.1 101).
type Named struct {
	ID    string   `json:"id"`
	Names []string `json:"names,omitempty"`
}
