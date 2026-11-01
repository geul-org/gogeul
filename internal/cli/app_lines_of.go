//ff:func feature=cli type=helper control=sequence
//ff:what app.linesOf — a document's records as notation lines, laid out by all of the data folder's codec specs
package cli

// linesOf is linesWith over every spec of the data folder (spec), for commands
// that walk documents; show reads only what its document needs (docSpec).
func (a *app) linesOf(b []byte) (docLines, error) {
	return linesWith(b, a.spec())
}
