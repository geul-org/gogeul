//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what packSources — resolver sources with the comma-separated GEULPACK items before the GEULPROXY ones
package cli

import (
	"strings"

	"github.com/geul-org/gogeul/pkg/resolve"
)

// packSources puts the GEULPACK items before cfg.Proxy. GEULPROXY items need
// no help: resolve finds out how each place stores its documents.
func packSources(cfg resolve.Config, packs string) (resolve.Config, error) {
	var pre []resolve.Entry
	for _, item := range strings.Split(packs, ",") {
		if item = strings.TrimSpace(item); item == "" {
			continue
		}
		e, err := packEntry(item)
		if err != nil {
			return cfg, err
		}
		pre = append(pre, e)
	}
	cfg.Proxy = append(pre, cfg.Proxy...)
	return cfg, nil
}
