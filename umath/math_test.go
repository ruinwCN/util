package umath

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax(t *testing.T) {
	assert.Equal(t, Max(1, 2), 2)
	assert.Equal(t, Max(1.5, 2.3), 2.3)
	assert.Equal(t, Max[int](1, 2), 2)
	assert.Equal(t, Max[int8](1, 2), int8(2))
	assert.Equal(t, Max[int16](1, 2), int16(2))
	assert.Equal(t, Max[int32](-1, 2), int32(2))
	assert.Equal(t, Max[int64](1, 2), int64(2))
	assert.Equal(t, Max[uint](1, 2), uint(2))
	assert.Equal(t, Max[uint8](1, 2), uint8(2))
	assert.Equal(t, Max[uint16](1, 2), uint16(2))
	assert.Equal(t, Max[uint32](1, 2), uint32(2))
	assert.Equal(t, Max[uint64](1, 2), uint64(2))
	assert.Equal(t, Max[float32](1, 2), float32(2.0))
	assert.Equal(t, Max[float64](1, 2), 2.0)
}

func TestMin(t *testing.T) {
	assert.Equal(t, Min(1, 2), 1)
	assert.Equal(t, Min(1.5, 2.3), 1.5)
	assert.Equal(t, Min[int](1, 2), 1)
	assert.Equal(t, Min[int8](1, 2), int8(1))
	assert.Equal(t, Min[int16](1, 2), int16(1))
	assert.Equal(t, Min[int32](-1, 2), int32(-1))
	assert.Equal(t, Min[int64](1, 2), int64(1))
	assert.Equal(t, Min[uint](1, 2), uint(1))
	assert.Equal(t, Min[uint8](1, 2), uint8(1))
	assert.Equal(t, Min[uint16](1, 2), uint16(1))
	assert.Equal(t, Min[uint32](1, 2), uint32(1))
	assert.Equal(t, Min[uint64](1, 2), uint64(1))
	assert.Equal(t, Min[float32](1, 2), float32(1.0))
	assert.Equal(t, Min[float64](1, 2), 1.0)
}
