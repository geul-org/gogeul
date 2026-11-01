//ff:func feature=cli type=helper control=sequence
//ff:what harness.app — test fixture: an app with the harness settings loaded and output in out
package cli

import (
	"io"

	"github.com/geul-org/gogeul/internal/config"
)

// app is what run builds, for calling one method directly.
func (h *harness) app(out io.Writer) *app {
	a := &app{env: config.Env{Getenv: func(k string) string { return h.env[k] }, Cwd: h.cwd, ConfigDir: h.conf, CacheDir: h.conf}, out: out}
	a.cfg, _ = config.Load(a.env)
	return a
}
