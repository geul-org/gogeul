//ff:type feature=resolve type=resolver
//ff:what Dir — a Source and publish target over an archive: `<Root>/.well-known/geul/` as a directory tree, any Store, or a Store over the loose files under Root
package resolve

// Dir is an archive laid out like the public address space. Store nil means
// the directory tree `<Root>/.well-known/geul/<path>` (dumps, mirrors and
// caches share it). A Store (pack sets) comes first; with a Root too, a path
// it lacks is read as a loose file of that tree, so a data folder may keep
// some Issuers as sets and others as files. Writes go to the Store.
type Dir struct {
	Root  string
	Store Store
}
