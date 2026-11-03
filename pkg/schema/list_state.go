//ff:type feature=schema type=model
//ff:what listState — what a version list has said so far: its Issuer, document and version state
package schema

// listState is the state of a version list being checked: the Issuer and the
// document it lists (canonical hex) and that document's version state.
type listState struct {
	issuer string
	doc    string
	docs   map[string]*docVersions
}
