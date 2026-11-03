//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what TestDocOfID — a 16-digit lowercase document ID as 4 words; names and other widths are ErrNotFound
package resolve

import (
	"errors"
	"testing"
)

func TestDocOfID(t *testing.T) {
	if ws, err := docOfID(nil, "0x0001000200030004"); err != nil || len(ws) != 4 || ws[3] != 4 {
		t.Fatal(ws, err)
	}
	for _, s := range []string{"seven", "0x00000980", "0x000100020003000A"} {
		if _, err := docOfID(nil, s); !errors.Is(err, ErrNotFound) {
			t.Fatal(s, err)
		}
	}
}
