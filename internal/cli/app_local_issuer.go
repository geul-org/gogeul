//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what app.localIssuer — an address whose Issuer part names a local key, with that key's Issuer ID in its place
package cli

import (
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
)

// localIssuer replaces an Issuer part of s that names a local key (localKeys)
// by that 16-word Issuer's ID.
func (a *app) localIssuer(s string) (string, error) {
	is, rest, _ := strings.Cut(s, "/")
	keys, err := a.localKeys()
	if err != nil {
		return "", err
	}
	for _, k := range keys {
		if k.Name == is {
			s = core.FormatID(k.Issuer) + "/" + rest
		}
	}
	return s, nil
}
