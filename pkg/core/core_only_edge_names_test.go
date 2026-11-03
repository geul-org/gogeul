//ff:func feature=core type=helper control=sequence
//ff:what coreUnnamedEdge — Edge names the vector holds but core leaves to pkg/toulmin or has no constant for
package core

// coreUnnamedEdge reports whether the Edge name is one core does not name: the
// TOULMIN family (named by pkg/toulmin, Codec 0x60~0x7B) and the three grammar
// codes gogeul has no constant for (SUPERSEDE-CHAIN, VERDICT, EVIDENCE;
// migration.md §3.5).
func coreUnnamedEdge(codec uint16) bool {
	return codec >= 0x60 && codec <= 0x7B || codec >= 0x80 && codec <= 0x94
}
