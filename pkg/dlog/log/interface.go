/*
 * File: interface.go
 * Date: Tuesday, 13th August 2024 9:39:18 PM
 * Author: Hy Nguyen
 * Copyright - 2024, HYSOFTWARE LLC
 */

package log

import (
	"context"
	"io"
	"time"
)

// IEvent represent a "logging event". It collects all different fields
// of the log line before it is published to log output
type IEvent interface {
	// Bool adds the key/value pair of a bool to the event
	// The same event is returned
	Bool(key string, val bool) IEvent

	// Bools adds the key/value pair of multi bools to the event
	// The same event is returned
	Bools(key string, val []bool) IEvent

	// Str adds the key/value pair of a string to the event
	// The same event is returned
	Str(key, val string) IEvent

	// Strs adds the key/value pair of multi strings to the event
	// The same event is returned
	Strs(key string, val []string) IEvent

	// Int adds the key/value pair of an int to the event
	// The same event is returned
	Int(key string, val int) IEvent

	// Ints adds the key/value pair of multi intsto the event
	// The same event is returned
	Ints(key string, val []int) IEvent

	// Int64 adds the key/value pair of an int64 to the event
	// The same event is returned
	Int64(key string, val int64) IEvent

	// Int64s adds the key/value pair of multi int64s to the event
	// The same event is returned
	Int64s(key string, val []int64) IEvent

	// Int32 adds the key/value pair of an int32 to the event
	// The same event is returned
	Int32(key string, val int32) IEvent

	// Int32s adds the key/value pair of multi int32s to the event
	// The same event is returned
	Int32s(key string, val []int32) IEvent

	// Int16 adds the key/value pair of an int16 to the event
	// The same event is returned
	Int16(key string, val int16) IEvent

	// Int16s adds the key/value pair of multi int16 to the event
	// The same event is returned
	Int16s(key string, val []int16) IEvent

	// Int8 adds the key/value pair of an int8 to the event
	// The same event is returned
	Int8(key string, val int8) IEvent

	// Int8s adds the key/value pair of multi int8s to the event
	// The same event is returned
	Int8s(key string, val []int8) IEvent

	// UInt adds the key/value pair of an uint to the event
	// The same event is returned
	UInt(key string, val uint) IEvent

	// UInts adds the key/value pair of multi uints to the event
	// The same event is returned
	UInts(key string, val []uint) IEvent

	// UInt64 adds the key/value pair of an uint64 to the event
	// The same event is returned
	UInt64(key string, val uint64) IEvent

	// UInt64s adds the key/value pair of multi uint64s to the event
	// The same event is returned
	UInt64s(key string, val []uint64) IEvent

	// UInt32 adds the key/value pair of  an uint32 to the event
	// The same event is returned
	UInt32(key string, val uint32) IEvent

	// UInt32s adds the key/value pair of multi uint32s to the event
	// The same event is returned
	UInt32s(key string, val []uint32) IEvent

	// UInt16 adds the key/value pair of and uint16 to the event
	// The same event is returned
	UInt16(key string, val uint16) IEvent

	// UInt16s adds the key/value pair of multi uint16s to the event
	// The same event is returned
	UInt16s(key string, val []uint16) IEvent

	// UInt8 adds the key/value pair of an uint 8 to the event
	// The same event is returned
	UInt8(key string, val uint8) IEvent

	// UInt8s adds the key/value pair of multi uint8s to the event
	// The same event is returned
	UInt8s(key string, val []uint8) IEvent

	// Float64 adds the key/value pair of a float64 to the event
	// The same event is returned
	Float64(key string, val float64) IEvent

	// Float64s adds the key/value pair of multi float64s to the event
	// The same event is returned
	Float64s(key string, val []float64) IEvent

	// Float32 adds the key/value pair of a float32 to the event
	// The same event is returned
	Float32(key string, val float32) IEvent

	// Float32s adds the key/value pair of multi float32s to the event
	// The same event is returned
	Float32s(key string, val []float32) IEvent

	// AnErr adds the key/value pair of and error to the event
	// The same event is returned
	AnErr(key string, val error) IEvent

	// Errs adds the key/value pair of multi errors to the event
	// The same event is returned
	Errs(key string, val []error) IEvent

	// Bytes adds the key/value pair of data bytes to the event
	// The same event is returned
	Bytes(key string, val []byte) IEvent

	// Hex adds the key/value pair of hex bytes to the event
	// The same event is returned
	Hex(key string, val []byte) IEvent

	// RawJSON adds the key/value pair of raw JSON data to the event
	// The same event is returned
	RawJSON(key string, val []byte) IEvent

	// Dur adds the key/value pair of a duration to the event
	// The same event is returned
	Dur(key string, val time.Duration) IEvent

	// Durs adds the key/value pair of multi durations to the event
	// The same event is returned
	Durs(key string, val []time.Duration) IEvent
}

// Prop is the function type that add property field to IEvent
type Prop func(evt IEvent)

// ILogger is the interface of the logger, the main mechanism
// to do logging
type ILogger interface {
	// With creates child logger and return the associated IContext
	With() IContext

	// Context pairs the ILogger with the given Context.
	// The ILogger is set as a value in the given Context
	// The updated Context is then returned
	Context(ctx context.Context) context.Context

	// Level updates the level of the ILogger and returns
	// the ILogger
	Level(l Level) ILogger

	// Output replaces the output writer with the given one
	// and returns the ILogger
	Output(w io.Writer) ILogger

	// DebugMsg logs the given message with DEBUG level
	DebugMsg(msg string)

	// InfoMsg logs the given message with INFO level
	InfoMsg(msg string)

	// WarnMsg logs the given message with WARN level
	WarnMsg(msg string)

	// ErrorMsg logs the given message with ERROR level
	ErrorMsg(msg string)

	// PanicMsg logs the given message with PANIC level
	PanicMsg(msg string)

	// Debug logs the given message with DEBUG level along with any
	// provided properties
	Debug(msg string, props ...Prop)

	// Info logs the given message with INFO level along with any
	// provided properties
	Info(msg string, props ...Prop)

	// Warn logs the given message with WARN level along with any
	// provided properties
	Warn(msg string, props ...Prop)

	// Error logs the given message with ERROR level along with any
	// provided properties
	Error(msg string, props ...Prop)

	// Panic logs the given message with PANIC level along with any
	// provided properties
	Panic(msg string, props ...Prop)
}

// IContext provides the mechanism to create child logger
// with preset properties
type IContext interface {
	// Logger creates the child logger tied to this IContext
	Logger() ILogger

	// Bool presets key/val pair of and return the
	// updated IContext
	Bool(key string, val bool) IContext

	// Bools presets key/val pair of multi bools and return the
	// updated IContext
	Bools(key string, val []bool) IContext

	// Str presets key/val pair of a string and return the
	// updated IContext
	Str(key, val string) IContext

	// Strs presets key/val pair of multi strings and return the
	// updated IContext
	Strs(key string, val []string) IContext

	// Int presets key/val pair of an int and return the
	// updated IContext
	Int(key string, val int) IContext

	// Ints presets key/val pair of multi ints and return the
	// updated IContext
	Ints(key string, val []int) IContext

	// Int64 presets key/val pair of an int64 and return the
	// updated IContext
	Int64(key string, val int64) IContext

	// Int64s presets key/val pair of multi int64s and return the
	// updated IContext
	Int64s(key string, val []int64) IContext

	// Int32 presets key/val pair of an int32 and return the
	// updated IContext
	Int32(key string, val int32) IContext

	// Int32s presets key/val pair of multi int32s and return the
	// updated IContext
	Int32s(key string, val []int32) IContext

	// Int16 presets key/val pair of an int16 and return the
	// updated IContext
	Int16(key string, val int16) IContext

	// Int16s presets key/val pair of multi int16s and return the
	// updated IContext
	Int16s(key string, val []int16) IContext

	// Int8 presets key/val pair of an int8 and return the
	// updated IContext
	Int8(key string, val int8) IContext

	// Int8s presets key/val pair of multi int8s and return the
	// updated IContext
	Int8s(key string, val []int8) IContext

	// UInt presets key/val pair of an uint and return the
	// updated IContext
	UInt(key string, val uint) IContext

	// UInts presets key/val pair of multi uints and return the
	// updated IContext
	UInts(key string, val []uint) IContext

	// UInt64 presets key/val pair of an uint64 and return the
	// updated IContext
	UInt64(key string, val uint64) IContext

	// UInt64s presets key/val pair of multi uint64s and return the
	// updated IContext
	UInt64s(key string, val []uint64) IContext

	// Uint32  presets key/val pair of an uint32 and return the
	// updated IContext
	UInt32(key string, val uint32) IContext

	// Uint32s presets key/val pair of multi uint32s and return the
	// updated IContext
	UInt32s(key string, val []uint32) IContext

	// UInt16 presets key/val pair of an uint16 and return the
	// updated IContext
	UInt16(key string, val uint16) IContext

	// UInt16s presets key/val pair of multi uint16s and return the
	// updated IContext
	UInt16s(key string, val []uint16) IContext

	// UInt8 presets key/val pair of an uint8 and return the
	// updated IContext
	UInt8(key string, val uint8) IContext

	// UInt8s presets key/val pair of multi uint8s and return the
	// updated IContext
	UInt8s(key string, val []uint8) IContext

	// Float64 presets key/val pair of a float64and return the
	// updated IContext
	Float64(key string, val float64) IContext

	// Float64s presets key/val pair multi float64s of and return the
	// updated IContext
	Float64s(key string, val []float64) IContext

	// Float32 presets key/val pair of a float32 and return the
	// updated IContext
	Float32(key string, val float32) IContext

	// Float32s presets key/val pair of multi float32s and return the
	// updated IContext
	Float32s(key string, val []float32) IContext

	// AnErr presets key/val pair of an error and return the
	// updated IContext
	AnErr(key string, val error) IContext

	// Errs presets key/val pair of errors and return the
	// updated IContext
	Errs(key string, val []error) IContext

	// Bytes presets key/val pair of bytes and return the
	// updated IContext
	Bytes(key string, val []byte) IContext

	// Hex presets key/val pair of a hex data and return the
	// updated IContext
	Hex(key string, val []byte) IContext

	// RawJSON presets key/val pair of raw JSON data and return the
	// updated IContext
	RawJSON(key string, val []byte) IContext

	// Dur presets key/val pair of a duration and return the
	// updated IContext
	Dur(key string, val time.Duration) IContext

	// Durs presets key/val pair of multi durations and return the
	// updated IContext
	Durs(key string, val []time.Duration) IContext
}

// CtxKey is the context key of the Logger value in the context.Context
type CtxKey struct{}
