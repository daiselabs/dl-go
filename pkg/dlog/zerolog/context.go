/*
 * File: context.go
 * Date: Wednesday, 21st August 2024 10:32:53 PM
 * Author: Hy Nguyen
 * Copyright - 2024, HYSOFTWARE LLC
 */

package zerolog

import (
	"time"

	"github.com/daiselabs/dl-go/pkg/dlog/log"
	"github.com/rs/zerolog"
)

type zerologContext struct {
	context *zerolog.Context
	level   log.Level
}

func (zlc *zerologContext) Logger() log.ILogger {
	l := zlc.context.Logger()
	return &zerologLogger{zl: &l, level: zlc.level}
}

func (zlc *zerologContext) Level() log.Level {
	return zlc.level
}

func (zlc *zerologContext) Bool(key string, val bool) log.IContext {
	c := zlc.context.Bool(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) Bools(key string, val []bool) log.IContext {
	c := zlc.context.Bools(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) Str(key, val string) log.IContext {
	c := zlc.context.Str(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) Strs(key string, val []string) log.IContext {
	c := zlc.context.Strs(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) Int(key string, val int) log.IContext {
	c := zlc.context.Int(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) Ints(key string, val []int) log.IContext {
	c := zlc.context.Ints(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) Int64(key string, val int64) log.IContext {
	c := zlc.context.Int64(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) Int64s(key string, val []int64) log.IContext {
	c := zlc.context.Ints64(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) Int32(key string, val int32) log.IContext {
	c := zlc.context.Int32(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) Int32s(key string, val []int32) log.IContext {
	c := zlc.context.Ints32(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) Int16(key string, val int16) log.IContext {
	c := zlc.context.Int16(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) Int16s(key string, val []int16) log.IContext {
	c := zlc.context.Ints16(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) Int8(key string, val int8) log.IContext {
	c := zlc.context.Int8(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) Int8s(key string, val []int8) log.IContext {
	c := zlc.context.Ints8(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) UInt(key string, val uint) log.IContext {
	c := zlc.context.Uint(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) UInts(key string, val []uint) log.IContext {
	c := zlc.context.Uints(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) UInt64(key string, val uint64) log.IContext {
	c := zlc.context.Uint64(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) UInt64s(key string, val []uint64) log.IContext {
	c := zlc.context.Uints64(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) UInt32(key string, val uint32) log.IContext {
	c := zlc.context.Uint32(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) UInt32s(key string, val []uint32) log.IContext {
	c := zlc.context.Uints32(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) UInt16(key string, val uint16) log.IContext {
	c := zlc.context.Uint16(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) UInt16s(key string, val []uint16) log.IContext {
	c := zlc.context.Uints16(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) UInt8(key string, val uint8) log.IContext {
	c := zlc.context.Uint8(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) UInt8s(key string, val []uint8) log.IContext {
	c := zlc.context.Uints8(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) Float64(key string, val float64) log.IContext {
	c := zlc.context.Float64(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) Float64s(key string, val []float64) log.IContext {
	c := zlc.context.Floats64(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) Float32(key string, val float32) log.IContext {
	c := zlc.context.Float32(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) Float32s(key string, val []float32) log.IContext {
	c := zlc.context.Floats32(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) AnErr(key string, val error) log.IContext {
	c := zlc.context.AnErr(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) Errs(key string, val []error) log.IContext {
	c := zlc.context.Errs(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) Bytes(key string, val []byte) log.IContext {
	c := zlc.context.Bytes(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) Hex(key string, val []byte) log.IContext {
	c := zlc.context.Hex(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) RawJSON(key string, val []byte) log.IContext {
	c := zlc.context.RawJSON(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) Dur(key string, val time.Duration) log.IContext {
	c := zlc.context.Dur(key, val)
	zlc.context = &c
	return zlc
}

func (zlc *zerologContext) Durs(key string, val []time.Duration) log.IContext {
	c := zlc.context.Durs(key, val)
	zlc.context = &c
	return zlc
}
