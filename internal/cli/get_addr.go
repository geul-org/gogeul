//ff:func feature=cli type=command control=sequence
//ff:what getAddr — parse a 공표 주소 and fetch its bytes (geul get)
package cli

import (
	"context"

	"github.com/geul-org/gogeul/pkg/resolve"
)

// getAddr is geul get without the output: resolve.ParseAddr then
// resolve.Fetch. A form outside the ten or a missing address wraps
// resolve.ErrNotFound.
func getAddr(ctx context.Context, s string) ([]byte, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	a, err := resolve.ParseAddr(s)
	if err != nil {
		return nil, err
	}
	return resolve.Fetch(ctx, a)
}
