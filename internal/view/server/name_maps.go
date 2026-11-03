//ff:type feature=server type=model
//ff:what nameMaps — an Issuer's document names both ways: name → ID and ID → name
package server

// nameMaps is one parse of an Issuer's document list (`<Issuer>/docs/`). Callers only read it.
type nameMaps struct {
	byName map[string]string
	byID   map[string]string
}
