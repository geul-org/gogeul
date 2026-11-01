//ff:type feature=pack type=model
//ff:what crashCodec — a Codec that passes to another for a number of calls, then fails: a Close stopped at that compression
package pack

// crashCodec compresses with inner while *left > 0 (counting down), then
// fails, standing for a process stopped at that point of Close.
type crashCodec struct {
	inner Codec
	left  *int
}
