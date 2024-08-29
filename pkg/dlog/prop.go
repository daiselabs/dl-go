/*
 * File: prop.go
 * Date: Wednesday, 14th August 2024 9:12:10 PM
 * Author: Hy Nguyen
 * Copyright - 2024, HYSOFTWARE LLC
 */
package dlog

import (
	"time"

	"github.com/daiselabs/dl-go/pkg/dlog/log"
)

var (
	rootLevel            log.Level = log.InfoLevel
	rootTimestampEnabled bool      = false
)

// SetRootLevel sets the default level of the root logger
func SetRootLevel(l log.Level) {
	rootLevel = l
}

// GetRootLevel get the current default level of the root logger
func GetRootLevel() log.Level {
	return rootLevel
}

// SetRootTimeStampEnabled enables/disables the timestamp field
// of the root logger
func SetRootTimeStampEnabled(f bool) {
	rootTimestampEnabled = f
}

// GetRootTimeStampEnabled gets the flag that enables/disables
// the timestamp field of the root logger
func GetRootTimeStampEnabled() bool {
	return rootTimestampEnabled
}

// Ints creates a property setter for key/value pair of multi ints
func Ints(key string, val []int) log.Prop {
	return func(evt log.IEvent) {
		evt.Ints(key, val)
	}
}

// Int8s creates a property setter for key/value pair of multi int8s
func Int8s(key string, val []int8) log.Prop {
	return func(evt log.IEvent) {
		evt.Int8s(key, val)
	}
}

// Int16s creates a property setter for key/value pair of multi int16s
func Int16s(key string, val []int16) log.Prop {
	return func(evt log.IEvent) {
		evt.Int16s(key, val)
	}
}

// Int32s creates a property setter for key/value pair of multi int32s
func Int32s(key string, val []int32) log.Prop {
	return func(evt log.IEvent) {
		evt.Int32s(key, val)
	}
}

// Int64s creates a property setter for key/value pair of multi int64s
func Int64s(key string, val []int64) log.Prop {
	return func(evt log.IEvent) {
		evt.Int64s(key, val)
	}
}

// UInts creates a property setter for key/value pair of multi uints
func UInts(key string, val []uint) log.Prop {
	return func(evt log.IEvent) {
		evt.UInts(key, val)
	}
}

// UInt8s creates a property setter for key/value pair of multi uint8s
func UInt8s(key string, val []uint8) log.Prop {
	return func(evt log.IEvent) {
		evt.UInt8s(key, val)
	}
}

// UInt16s creates a property setter for key/value pair of multi uint16s
func UInt16s(key string, val []uint16) log.Prop {
	return func(evt log.IEvent) {
		evt.UInt16s(key, val)
	}
}

// UInt32s creates a property setter for key/value pair of multi uint32s
func UInt32s(key string, val []uint32) log.Prop {
	return func(evt log.IEvent) {
		evt.UInt32s(key, val)
	}
}

// UInt64s creates a property setter for key/value pair of multi uint64s
func UInt64s(key string, val []uint64) log.Prop {
	return func(evt log.IEvent) {
		evt.UInt64s(key, val)
	}
}

// Int creates a property setter for key/value pair of an int
func Int(key string, val int) log.Prop {
	return func(evt log.IEvent) {
		evt.Int(key, val)
	}
}

// Int8 creates a property setter for key/value pair of an int9
func Int8(key string, val int8) log.Prop {
	return func(evt log.IEvent) {
		evt.Int8(key, val)
	}
}

// Int16 creates a property setter for key/value pair of an int16
func Int16(key string, val int16) log.Prop {
	return func(evt log.IEvent) {
		evt.Int16(key, val)
	}
}

// Int32 creates a property setter for key/value pair of int32
func Int32(key string, val int32) log.Prop {
	return func(evt log.IEvent) {
		evt.Int32(key, val)
	}
}

// Int64 creates a property setter for key/value pair of an int64
func Int64(key string, val int64) log.Prop {
	return func(evt log.IEvent) {
		evt.Int64(key, val)
	}
}

// UInt creates a property setter for key/value pair of an uint
func UInt(key string, val uint) log.Prop {
	return func(evt log.IEvent) {
		evt.UInt(key, val)
	}
}

// UInt8 creates a property setter for key/value pair of an uint8
func UInt8(key string, val uint8) log.Prop {
	return func(evt log.IEvent) {
		evt.UInt8(key, val)
	}
}

// UInt16 creates a property setter for key/value pair of an uint16
func UInt16(key string, val uint16) log.Prop {
	return func(evt log.IEvent) {
		evt.UInt16(key, val)
	}
}

// UInt32 creates a property setter for key/value pair of an uint32
func UInt32(key string, val uint32) log.Prop {
	return func(evt log.IEvent) {
		evt.UInt32(key, val)
	}
}

// UInt64 creates a property setter for key/value pair of an uint64
func UInt64(key string, val uint64) log.Prop {
	return func(evt log.IEvent) {
		evt.UInt64(key, val)
	}
}

// Float32 creates a property setter for key/value pair of a float32
func Float32(key string, val float32) log.Prop {
	return func(evt log.IEvent) {
		evt.Float32(key, val)
	}
}

// Float64 creates a property setter for key/value pair of a float64
func Float64(key string, val float64) log.Prop {
	return func(evt log.IEvent) {
		evt.Float64(key, val)
	}
}

// Float32s creates a property setter for key/value pair of multi float32s
func Float32s(key string, val []float32) log.Prop {
	return func(evt log.IEvent) {
		evt.Float32s(key, val)
	}
}

// Float64s creates a property setter for key/value pair of multi float64s
func Float64s(key string, val []float64) log.Prop {
	return func(evt log.IEvent) {
		evt.Float64s(key, val)
	}
}

// AnErr creates a property setter for key/value pair of an error
func AnErr(key string, val error) log.Prop {
	return func(evt log.IEvent) {
		evt.AnErr(key, val)
	}
}

// Errs creates a property setter for key/value pair of multi errors
func Errs(key string, val []error) log.Prop {
	return func(evt log.IEvent) {
		evt.Errs(key, val)
	}
}

// Str creates a property setter for key/value pair of a string
func Str(key string, val string) log.Prop {
	return func(evt log.IEvent) {
		evt.Str(key, val)
	}
}

// Strs creates a property setter for key/value pair of
func Strs(key string, val []string) log.Prop {
	return func(evt log.IEvent) {
		evt.Strs(key, val)
	}
}

// RawJSON creates a property setter for key/value pair of raw JSON data
func RawJSON(key string, val []byte) log.Prop {
	return func(evt log.IEvent) {
		evt.RawJSON(key, val)
	}
}

// Bytes creates a property setter for key/value pair of bytes
func Bytes(key string, val []byte) log.Prop {
	return func(evt log.IEvent) {
		evt.Bytes(key, val)
	}
}

// Hex creates a property setter for key/value pair of hex data
func Hex(key string, val []byte) log.Prop {
	return func(evt log.IEvent) {
		evt.Hex(key, val)
	}
}

// Dur creates a property setter for key/value pair of a duration
func Dur(key string, val time.Duration) log.Prop {
	return func(evt log.IEvent) {
		evt.Dur(key, val)
	}
}

// Durs creates a property setter for key/value pair of multi duration
func Durs(key string, val []time.Duration) log.Prop {
	return func(evt log.IEvent) {
		evt.Durs(key, val)
	}
}
