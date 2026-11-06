//ff:type feature=core type=model
//ff:what layoutTable — the plan2 layout (testdata/migration1/layout2.md): Kind and title of each 32-code section, name of each assigned code
package core

// layoutTable is the section table read from layout2.md: kind and title of
// the 16 sections (w>>9) and the table name of every assigned 9-bit code
// (w>>4); an unassigned code has no entry.
type layoutTable struct {
	kind  [16]Kind
	title [16]string
	names map[uint16]string
}
