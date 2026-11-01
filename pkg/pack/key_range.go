//ff:type feature=pack type=model
//ff:what keyRange — the keys of one directory: the layers holding them, [lo, hi) and how a key spells its path
package pack

// keyRange is where a directory's paths live in the key spaces (§3.4): keys
// lo ≤ key < hi of layers (hi "" is unbounded), each named by path.
type keyRange struct {
	layers []layer
	lo, hi string
	path   func(string) string
}
