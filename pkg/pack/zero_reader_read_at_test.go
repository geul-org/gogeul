//ff:func feature=pack type=helper control=sequence
//ff:what zeroReader.ReadAt — fill b with zeros
package pack

// ReadAt zeroes b and reports it full.
func (zeroReader) ReadAt(b []byte, off int64) (int, error) {
	clear(b)
	return len(b), nil
}
