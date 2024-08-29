/*
 * File: level.go
 * Date: Tuesday, 20th August 2024 9:54:04 PM
 * Author: Hy Nguyen
 * Copyright - 2024, HYSOFTWARE LLC
 */

package zerolog

import (
	"github.com/daiselabs/dl-go/pkg/dlog/log"
	"github.com/rs/zerolog"
)

func fromZerologLevel(zll zerolog.Level) log.Level {
	switch zll {
	case zerolog.Disabled:
		return log.Disabled
	case zerolog.DebugLevel:
		return log.DebugLevel
	case zerolog.InfoLevel:
		return log.InfoLevel
	case zerolog.WarnLevel:
		return log.WarnLevel
	case zerolog.ErrorLevel:
		return log.ErrorLevel
	case zerolog.FatalLevel:
		return log.FatalLevel
	case zerolog.PanicLevel:
		return log.PanicLevel
	case zerolog.NoLevel:
		return log.NoLevel
	case zerolog.TraceLevel:
		return log.TraceLevel
	}
	return log.NoLevel
}

func toZerologLevel(l log.Level) zerolog.Level {
	switch l {
	case log.Disabled:
		return zerolog.Disabled
	case log.DebugLevel:
		return zerolog.DebugLevel
	case log.InfoLevel:
		return zerolog.InfoLevel
	case log.WarnLevel:
		return zerolog.WarnLevel
	case log.ErrorLevel:
		return zerolog.ErrorLevel
	case log.FatalLevel:
		return zerolog.FatalLevel
	case log.PanicLevel:
		return zerolog.PanicLevel
	case log.NoLevel:
		return zerolog.NoLevel
	case log.TraceLevel:
		return zerolog.TraceLevel
	}
	return zerolog.NoLevel
}
