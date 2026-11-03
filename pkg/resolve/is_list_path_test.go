//ff:func feature=resolve type=rule control=sequence
//ff:what TestIsListPath — versions.geul paths only
package resolve

import (
	"testing"
)

func TestIsListPath(t *testing.T) {
	if !isListPath("0x00000000/docs/0x0000000000000001/versions.geul") || isListPath("0x00000000/docs/0x0000000000000001/0x00000001.geul") || isListPath("versions.geul") {
		t.Fatal("wrong")
	}
}
