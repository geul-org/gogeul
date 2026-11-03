//ff:func feature=resolve type=helper control=selection
//ff:what sourceOfBase — the Source an address base reads from: file://<dir> (files or pack sets), https://<host> (HTTP, Range on pack sets)
package resolve

import "fmt"

// sourceOfBase opens the place an Addr's Base names, as a GEULPROXY item of
// that spelling does (parseEntry): any storage form is found when read.
func sourceOfBase(base string) (Source, error) {
	if base == "" {
		return nil, fmt.Errorf("geul/resolve: an address without file:// or https:// has no place to read")
	}
	e, err := parseEntry(base)
	if err != nil {
		return nil, err
	}
	return e.Source, nil
}
