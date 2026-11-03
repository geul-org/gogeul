//ff:type feature=build type=model
//ff:what indexEntry — one link of a directory listing
package build

// indexEntry is one entry of a directory listing: the name shown and its URL.
// ID is the Issuer ID when Name shows the Issuer's bound name (the Issuer list)
// instead of the ID; the URL keeps the ID path.
type indexEntry struct {
	Name string
	URL  string
	ID   string `json:",omitempty"`
}
