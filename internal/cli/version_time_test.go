//ff:func feature=cli type=helper control=sequence
//ff:what TestVersionTime — UTC time; zero is -
package cli

import (
	"testing"
	"time"
)

func TestVersionTime(t *testing.T) {
	if versionTime(time.Unix(0, 0)) != "1970-01-01T00:00:00Z" || versionTime(time.Time{}) != "-" {
		t.Fatal("versionTime")
	}
}
