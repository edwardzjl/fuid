package uuid

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/mariuszs/friendlyid-go/friendlyid"
)

type UUID struct {
	value uuid.UUID
}

var Nil = UUID{uuid.Nil}

func New() UUID {
	value := uuid.New()
	return UUID{value: value}
}

func Parse(raw string) (UUID, error) {
	decoded, err := friendlyid.Decode(raw)
	if err != nil {
		return Nil, err
	}
	value, err := uuid.Parse(decoded)
	if err != nil {
		return Nil, err
	}
	return UUID{value: value}, nil
}

func MustParse(s string) UUID {
	uuid, err := Parse(s)
	if err != nil {
		panic(`uuid: Parse(` + s + `): ` + err.Error())
	}
	return uuid
}

func (fid UUID) String() string {
	uuidStr := fid.value.String()
	res, err := friendlyid.Encode(uuidStr)
	if err != nil {
		return ""
	}
	return res
}

func (fid UUID) MarshalJSON() ([]byte, error) {
	return json.Marshal(fid.String())
}

func (fid *UUID) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	id, err := Parse(value)
	if err != nil {
		return err
	}
	*fid = id
	return nil
}

func (fid UUID) Value() (driver.Value, error) {
	return fid.value.Value()
}

func (fid *UUID) Scan(src interface{}) error {
	var value uuid.UUID
	err := fid.value.Scan(src)
	if err != nil {
		return err
	}
	fid = &UUID{value: value}
	return nil
}

// URN returns the RFC 2141 URN form of uuid,
// urn:uuid:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx,  or "" if uuid is invalid.
func (fid UUID) URN() string {
	return fid.value.URN()
}
