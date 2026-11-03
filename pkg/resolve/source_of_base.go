//ff:func feature=resolve type=helper control=sequence
//ff:what sourceOfBase — the Source an address base reads from: file://<dir> (files or pack sets), https://<host> (HTTP, Range on pack sets)
package resolve

import "fmt"

// sourceOfBase opens the place an Addr's Base names, as a GEULPROXY item of
// that spelling does (parseEntry): any storage form is found when read. A
// bare path names no place: `geul get` reads file:// and https:// addresses
// only, so it is ErrNotFound like every other address outside the forms
// (publishing.md).
func sourceOfBase(base string) (Source, error) {
	if base == "" {
		return nil, fmt.Errorf("%w: an address without file:// or https:// has no place to read", ErrNotFound)
	}
	e, err := parseEntry(base)
	if err != nil {
		return nil, err
	}
	return e.Source, nil
}
