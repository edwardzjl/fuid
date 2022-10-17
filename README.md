# FUID - A more human Friendly UUID

FUID extends [uuid](github.com/google/uuid) with [friendlyid](github.com/mariuszs/friendlyid-go/friendlyid) to provide a more human friendly uuid.

FUID wraps a `uuid`, and only modifies the string representation of a `uuid`, including `String()`, `Parse()`, `MarshalJson()` and `UnmarshalJson`.

## Install

```go
go get github.com/edwardzjl/fuid
```

## Usage

```go
import uuid "github.com/edwardzjl/fuid"

id := uuid.New()
```

## License

This project is under MIT License. See the [LICENSE](LICENSE) file for the full license text.
