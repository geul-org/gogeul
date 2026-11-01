//ff:func feature=pack type=helper control=sequence
//ff:what runName — the file name of H run n in the working folder
package pack

import "fmt"

// runName is "run-" and n in six digits (more when n needs them).
func runName(n int) string {
	return fmt.Sprintf("run-%06d", n)
}
