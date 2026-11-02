//ff:func feature=resolve type=resolver control=sequence
//ff:what TestLocalSite — objects and loose files both sit under `<root>/.well-known/geul/`
package resolve

import (
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
)

func TestLocalSite(t *testing.T) {
	s := localSite("/srv")
	wk := filepath.Join("/srv", ".well-known", "geul")
	if o, ok := s.objects.(pack.FileStore); !ok || o.Root != wk || s.dir != wk || s.loose != (Dir{Root: "/srv"}) || s.sets == nil {
		t.Fatalf("%+v", s)
	}
}
