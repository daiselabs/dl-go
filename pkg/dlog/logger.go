/*
 * File: logger.go
 * Date: Tuesday, 13th August 2024 9:43:09 PM
 * Author: Hy Nguyen
 * Copyright - 2024, HYSOFTWARE LLC
 */

package dlog

import (
	"context"
	"io"

	"github.com/daiselabs/dl-go/pkg/dlog/log"
	"github.com/daiselabs/dl-go/pkg/dlog/zerolog"
)

var (
	noopLogger log.ILogger = zerolog.NewZerologNoopLogger(log.Disabled)
)

// NewLogger creates a new Logger with a default provider
// root level
func NewLogger(w io.Writer) log.ILogger {
	return zerolog.NewZerologLogger(w, rootLevel)
}

// FromContext extract the log.ILogger stored in a given context.Context
// If none was stored, it'll return a "disabled" log.ILogger
func FromContext(ctx context.Context) log.ILogger {
	v := ctx.Value(log.CtxKey{})
	l, ok := v.(log.ILogger)
	if ok && l != nil {
		return l
	}
	return noopLogger
}
