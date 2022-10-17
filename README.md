# FUID - A more human Friendly UUID

FUID extends [uuid](github.com/google/uuid) with [friendlyid](github.com/mariuszs/friendlyid-go/friendlyid) to provide a more human friendly uuid.

FUID wraps a `uuid`, and only modifies the string representation of a `uuid`, including `String()`, `Parse()`, `MarshalJson()` and `UnmarshalJson`.
