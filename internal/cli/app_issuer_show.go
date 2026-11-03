//ff:func feature=cli type=command control=iteration dimension=1
//ff:what app.issuerShow — print one local Issuer by name or ID: ID, algorithm, public key
package cli

import (
	"crypto/ed25519"
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// issuerShow finds the local key named or identified by s.
func (a *app) issuerShow(s string) error {
	keys, err := a.localKeys()
	if err != nil {
		return err
	}
	for _, k := range keys {
		if k.Name == s || core.FormatID(k.Issuer) == s {
			fmt.Fprintf(a.out, "name %s\nissuer %s\nalgorithm Ed25519 (0x0001)\npublic-key 0x%X\n", k.Name, core.FormatID(k.Issuer), []byte(k.Key.Public().(ed25519.PublicKey)))
			return nil
		}
	}
	return fmt.Errorf("no local Issuer %q", s)
}
