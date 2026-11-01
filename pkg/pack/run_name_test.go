//ff:func feature=pack type=helper control=sequence
//ff:what TestRunName — six digits, more when needed
package pack

import "testing"

func TestRunName(t *testing.T) {
	if runName(7) != "run-000007" || runName(1234567) != "run-1234567" {
		t.Fatal(runName(7), runName(1234567))
	}
}
