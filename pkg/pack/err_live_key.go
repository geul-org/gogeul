//ff:type feature=pack type=model
//ff:what errLiveKey — the sentinel that stops a walk at the first live key
package pack

import "errors"

// errLiveKey ends Reader.walk as soon as anyLive has its answer; it never
// leaves anyLive.
var errLiveKey = errors.New("geul/pack: live key found")
