//ff:func feature=pack type=helper control=sequence
//ff:what notFound — ErrNotFound naming the path, so a resolver moves on to the next source
package pack

import (
	"fmt"
)

// notFound reports that this generation has no bytes at path.
func notFound(path string) error {
	return fmt.Errorf("%w: %s", ErrNotFound, path)
}
