//ff:type feature=core type=model
//ff:what vectorRow — one row of the migration1 reference vector: old TAG, new TAG, new name, old key
package core

// vectorRow is one line of testdata/migration1/vector.tsv. New is -1 when the
// old TAG has no TAG-only new value.
type vectorRow struct {
	Old    Word
	New    int
	Name   string
	OldKey string
}
