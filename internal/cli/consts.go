package cli

import "time"

// specTimeout bounds fetching the codec specs of one document: a source that
// does not answer leaves the records generic instead of hanging the command.
const specTimeout = 10 * time.Second
