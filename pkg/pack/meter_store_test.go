//ff:type feature=pack type=helper
//ff:what meterStore — test fixture: a Store that counts range reads, so a listing that walks the index shows
package pack

// meterStore counts ReadRange calls on s. Whole reads (latest.txt, manifest)
// happen once on open and are not counted.
type meterStore struct {
	Store
	reads int
}
