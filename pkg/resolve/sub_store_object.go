//ff:func feature=resolve type=helper control=sequence
//ff:what subStore.object — an object's name under the base
package resolve

// object prefixes object with the set's folder.
func (s subStore) object(object string) string {
	if s.name == "" {
		return object
	}
	return s.name + "/" + object
}
