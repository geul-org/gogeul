//ff:type feature=resolve type=model
//ff:what form — test fixture: one storage form of the same documents, read through a Dir
package resolve

// form is a way the fixture is stored and reached.
type form struct {
	name string
	dir  Dir
}
