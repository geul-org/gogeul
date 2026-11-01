//ff:func feature=cli type=helper control=sequence
//ff:what app.resolver — a Resolver with the data folder as the first source, then GEULPACK, GEULPROXY and GEULPRIVATE
package cli

import "github.com/geul-org/gogeul/pkg/resolve"

// resolver reads the source list from the environment and puts the data
// folder, when there is one, before it.
func (a *app) resolver() (*resolve.Resolver, error) {
	cfg, err := resolve.ConfigFromEnv(a.cfg.Get)
	if err == nil {
		cfg, err = packSources(cfg, a.cfg.Get("GEULPACK"))
	}
	if err != nil {
		return nil, err
	}
	if d, err := a.dir(); err == nil {
		cfg.Proxy = append([]resolve.Entry{{Name: "file://" + d.Root, Source: d}}, cfg.Proxy...)
	}
	return &resolve.Resolver{Config: cfg}, nil
}
