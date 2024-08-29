/*
 * File: event.go
 * Date: Monday, 19th August 2024 5:08:58 PM
 * Author: Hy Nguyen
 * Copyright - 2024, HYSOFTWARE LLC
 */

package zerolog

import (
	"time"

	"github.com/daiselabs/dl-go/pkg/dlog/log"
	"github.com/rs/zerolog"
)

type zlEvent struct {
	e *zerolog.Event
}

func (zle *zlEvent) Bool(key string, val bool) log.IEvent {
	zle.e.Bool(key, val)
	return zle
}

func (zle *zlEvent) Bools(key string, val []bool) log.IEvent {
	zle.e.Bools(key, val)
	return zle
}

func (zle *zlEvent) Str(key, val string) log.IEvent {
	zle.e.Str(key, val)
	return zle
}

func (zle *zlEvent) Strs(key string, val []string) log.IEvent {
	zle.e.Strs(key, val)
	return zle
}

func (zle *zlEvent) Int(key string, val int) log.IEvent {
	zle.e.Int(key, val)
	return zle
}

func (zle *zlEvent) Ints(key string, val []int) log.IEvent {
	zle.e.Ints(key, val)
	return zle
}

func (zle *zlEvent) Int64(key string, val int64) log.IEvent {
	zle.e.Int64(key, val)
	return zle
}

func (zle *zlEvent) Int64s(key string, val []int64) log.IEvent {
	zle.e.Ints64(key, val)
	return zle
}

func (zle *zlEvent) Int32(key string, val int32) log.IEvent {
	zle.e.Int32(key, val)
	return zle
}

func (zle *zlEvent) Int32s(key string, val []int32) log.IEvent {
	zle.e.Ints32(key, val)
	return zle
}

func (zle *zlEvent) Int16(key string, val int16) log.IEvent {
	zle.e.Int16(key, val)
	return zle
}

func (zle *zlEvent) Int16s(key string, val []int16) log.IEvent {
	zle.e.Ints16(key, val)
	return zle
}

func (zle *zlEvent) Int8(key string, val int8) log.IEvent {
	zle.e.Int8(key, val)
	return zle
}

func (zle *zlEvent) Int8s(key string, val []int8) log.IEvent {
	zle.e.Ints8(key, val)
	return zle
}

func (zle *zlEvent) UInt(key string, val uint) log.IEvent {
	zle.e.Uint(key, val)
	return zle
}

func (zle *zlEvent) UInts(key string, val []uint) log.IEvent {
	zle.e.Uints(key, val)
	return zle
}

func (zle *zlEvent) UInt64(key string, val uint64) log.IEvent {
	zle.e.Uint64(key, val)
	return zle
}

func (zle *zlEvent) UInt64s(key string, val []uint64) log.IEvent {
	zle.e.Uints64(key, val)
	return zle
}

func (zle *zlEvent) UInt32(key string, val uint32) log.IEvent {
	zle.e.Uint32(key, val)
	return zle
}

func (zle *zlEvent) UInt32s(key string, val []uint32) log.IEvent {
	zle.e.Uints32(key, val)
	return zle
}

func (zle *zlEvent) UInt16(key string, val uint16) log.IEvent {
	zle.e.Uint16(key, val)
	return zle
}

func (zle *zlEvent) UInt16s(key string, val []uint16) log.IEvent {
	zle.e.Uints16(key, val)
	return zle
}

func (zle *zlEvent) UInt8(key string, val uint8) log.IEvent {
	zle.e.Uint8(key, val)
	return zle
}

func (zle *zlEvent) UInt8s(key string, val []uint8) log.IEvent {
	zle.e.Uints8(key, val)
	return zle
}

func (zle *zlEvent) Float64(key string, val float64) log.IEvent {
	zle.e.Float64(key, val)
	return zle
}

func (zle *zlEvent) Float64s(key string, val []float64) log.IEvent {
	zle.e.Floats64(key, val)
	return zle
}

func (zle *zlEvent) Float32(key string, val float32) log.IEvent {
	zle.e.Float32(key, val)
	return zle
}

func (zle *zlEvent) Float32s(key string, val []float32) log.IEvent {
	zle.e.Floats32(key, val)
	return zle
}

func (zle *zlEvent) AnErr(key string, val error) log.IEvent {
	zle.e.AnErr(key, val)
	return zle
}

func (zle *zlEvent) Errs(key string, val []error) log.IEvent {
	zle.e.Errs(key, val)
	return zle
}

func (zle *zlEvent) Bytes(key string, val []byte) log.IEvent {
	zle.e.Bytes(key, val)
	return zle
}

func (zle *zlEvent) Hex(key string, val []byte) log.IEvent {
	zle.e.Hex(key, val)
	return zle
}

func (zle *zlEvent) RawJSON(key string, val []byte) log.IEvent {
	zle.e.RawJSON(key, val)
	return zle
}

func (zle *zlEvent) Dur(key string, val time.Duration) log.IEvent {
	zle.e.Dur(key, val)
	return zle
}

func (zle *zlEvent) Durs(key string, val []time.Duration) log.IEvent {
	zle.e.Durs(key, val)
	return zle
}
