//ff:func feature=pack type=helper control=sequence
//ff:what synthPath — the path of synth record i
package pack

import "fmt"

// synthPath names record i as synth does.
func synthPath(i int) string {
	return fmt.Sprintf("issuers/0x%08X.geul", i+1)
}
