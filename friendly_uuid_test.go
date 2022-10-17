package uuid

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestToString(t *testing.T) {
	id := New()
	str := id.String()
	fmt.Println(str)
}

func TestParse(t *testing.T) {
	id, err := Parse("1n559jchcg1vP4t5c9Mpfz")
	require.NoError(t, err)
	assert.NotEqual(t, Nil, id)
}

func TestMarshalJson(t *testing.T) {
	id, _ := Parse("1n559jchcg1vP4t5c9Mpfz")
	fid, err := json.Marshal(id)
	require.NoError(t, err)
	expected, _ := json.Marshal("1n559jchcg1vP4t5c9Mpfz")
	assert.Equal(t, expected, fid)
}

func TestUnmarshalJson(t *testing.T) {
	data := "\"1n559jchcg1vP4t5c9Mpfz\""
	var id UUID
	err := json.Unmarshal([]byte(data), &id)
	require.NoError(t, err)
	expected, _ := Parse("1n559jchcg1vP4t5c9Mpfz")
	assert.Equal(t, expected, id)
}

func TestValue(t *testing.T) {
	u := uuid.New()
	fid := UUID{Content: u}
	expected, _ := u.Value()
	actual, _ := fid.Value()
	assert.Equal(t, expected, actual)
}

func TestScan(t *testing.T) {
	fid := Nil
	err := fid.Scan(uuid.New().String())
	require.NoError(t, err)
}

func TestURN(t *testing.T) {
	fid := New()
	urn := fid.URN()
	fmt.Println(urn)
}
