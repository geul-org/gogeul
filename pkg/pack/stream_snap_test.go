//ff:type feature=pack type=model
//ff:what streamSnap — an open data pack as a checkpoint saves it: open object, inventory, paths
package pack

// streamSnap is the open data pack as a checkpoint would save it.
type streamSnap struct {
	open  streamOpen
	inv   []streamInv
	paths []string
}
