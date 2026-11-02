//ff:func feature=resolve type=resolver control=sequence
//ff:what overlay.WriteFile — Store: writes go to the stored side
package resolve

// WriteFile stores b at name.
func (o overlay) WriteFile(name string, b []byte) error {
	return o.stored.WriteFile(name, b)
}
