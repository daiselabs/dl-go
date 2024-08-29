/*
 * File: level.go
 * Date: Monday, 12th August 2024 9:44:05 PM
 * Author: Hy Nguyen
 * Copyright - 2024, HYSOFTWARE LLC
 */

package log

import (
	"fmt"
	"strings"
)

// Level is the logging level of a given log
type Level int8

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	PanicLevel
	NoLevel
	Disabled

	TraceLevel Level = -1
)

// String provides the string representation of the Level
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case FatalLevel:
		return "FATAL"
	case PanicLevel:
		return "PANIC"
	case NoLevel:
		return ""
	case Disabled:
		return "DISABLED"
	case TraceLevel:
		return "TRACE"
	}
	return fmt.Sprintf("LEVEL_%d", l)
}

// ParseLevelString parses a given string to get its Level value
func ParseLevelString(lstr string) Level {
	switch strings.ToUpper(strings.TrimSpace(lstr)) {
	case "DEBUG":
		return DebugLevel
	case "INFO":
		return InfoLevel
	case "WARN":
		return WarnLevel
	case "ERROR":
		return ErrorLevel
	case "FATAL":
		return FatalLevel
	case "PANIC":
		return PanicLevel
	case "TRACE":
		return TraceLevel
	}
	return NoLevel
}
