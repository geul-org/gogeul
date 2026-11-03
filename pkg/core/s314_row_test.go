//ff:type feature=core type=model
//ff:what s314Row — one row of testdata/migration1/s314.tsv (migration.md §3.14): a former GEUL ISSUER format and its new direct TAG
package core

// s314Row is a former GEUL Issuer ISSUER format (Kind, format number, name)
// and its direct code: the TAG with L=0 and the 32-bit codec ID.
type s314Row struct {
	Kind    string
	Format  uint32
	Name    string
	Tag     Word
	CodecID uint32
}
