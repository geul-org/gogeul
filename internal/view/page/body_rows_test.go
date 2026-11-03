//ff:func feature=page type=builder control=iteration dimension=1
//ff:what TestBodyRows — rows between CASE-BEGIN and CASE-END nest; a stray CASE-END never goes below 0
package page

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"

	"github.com/geul-org/gogeul/pkg/inspect"
	"github.com/geul-org/gogeul/pkg/toulmin"
)

func TestBodyRows(t *testing.T) {
	ctl := func(c toulmin.ControlCode) inspect.Entry {
		return inspect.Entry{Words: []core.Word{0x127C}, Data: toulmin.ControlPacket{Code: c}}
	}
	bare := inspect.Entry{Words: []core.Word{0x1260}}
	body := []inspect.Entry{ctl(toulmin.ControlCaseEnd), ctl(toulmin.ControlCaseBegin), bare, ctl(toulmin.ControlCaseEnd), bare}
	rows := bodyRows(body, nil)
	want := []int{0, 0, 1, 0, 0}
	for i, r := range rows {
		if r.Depth != want[i] {
			t.Fatalf("row %d depth %d want %d", i, r.Depth, want[i])
		}
	}
}
