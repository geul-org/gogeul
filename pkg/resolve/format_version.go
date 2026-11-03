//ff:func feature=resolve type=encoder control=sequence
//ff:what FormatVersion — a version ID as `0x` and 8 upper-case hex digits
package resolve

import "fmt"

// FormatVersion renders a 2-word version ID the way addresses write it
// (GEUL-ORG ID 표기): `0x` and 8 upper-case hex digits.
func FormatVersion(v uint32) string {
	return fmt.Sprintf("0x%08X", v)
}
