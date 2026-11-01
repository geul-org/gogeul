//ff:type feature=pack type=model
//ff:what zeroReader — an io.ReaderAt of endless zero bytes
package pack

// zeroReader reads zeros at any offset.
type zeroReader struct{}
