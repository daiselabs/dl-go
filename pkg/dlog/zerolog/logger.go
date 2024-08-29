/*
 * File: logger.go
 * Date: Tuesday, 13th August 2024 9:54:21 PM
 * Author: Hy Nguyen
 * Copyright - 2024, HYSOFTWARE LLC
 */

package zerolog

import (
	"context"
	"io"

	"github.com/daiselabs/dl-go/pkg/dlog/log"
	"github.com/rs/zerolog"
)

func init() {
	// Use the our level strings instead of zerolog provided ones
	zerolog.LevelFieldMarshalFunc = func(l zerolog.Level) string {
		return fromZerologLevel(l).String()
	}
}

type zerologLogger struct {
	zl    *zerolog.Logger
	level log.Level
}

func NewZerologLogger(w io.Writer, level log.Level) *zerologLogger {
	logger := zerolog.New(w).With().Logger()
	return &zerologLogger{zl: &logger, level: level}
}

func NewZerologNoopLogger(l log.Level) *zerologLogger {
	logger := zerolog.Nop()
	return &zerologLogger{zl: &logger, level: l}

}

func (lp *zerologLogger) shouldLog(l log.Level) bool {
	return l >= lp.level
}

func (lp *zerologLogger) With() log.IContext {
	c := lp.zl.With()
	return &zerologContext{context: &c, level: lp.level}
}

func (lp *zerologLogger) Context(ctx context.Context) context.Context {
	v := ctx.Value(log.CtxKey{})
	l, ok := v.(log.ILogger)
	if ok {
		if l == lp {
			return ctx
		}
	} else if lp.level == log.Disabled {
		return ctx
	}
	return context.WithValue(ctx, log.CtxKey{}, lp)
}

func (lp *zerologLogger) Level(l log.Level) log.ILogger {
	zl := lp.zl.Level(toZerologLevel(l))
	lp.zl = &zl
	lp.level = l
	return lp
}

func (lp *zerologLogger) Output(w io.Writer) log.ILogger {
	zl := lp.zl.Output(w)
	lp.zl = &zl
	return lp
}

func (lp *zerologLogger) InfoMsg(msg string) {
	if !lp.shouldLog(log.InfoLevel) {
		return
	}
	lp.zl.Info().Msg(msg)
}

func (lp *zerologLogger) DebugMsg(msg string) {
	if !lp.shouldLog(log.DebugLevel) {
		return
	}
	lp.zl.Debug().Msg(msg)
}

func (lp *zerologLogger) WarnMsg(msg string) {
	if !lp.shouldLog(log.WarnLevel) {
		return
	}
	lp.zl.Warn().Msg(msg)
}

func (lp *zerologLogger) ErrorMsg(msg string) {
	if !lp.shouldLog(log.ErrorLevel) {
		return
	}
	lp.zl.Error().Msg(msg)
}

func (lp *zerologLogger) PanicMsg(msg string) {
	if !lp.shouldLog(log.PanicLevel) {
		return
	}
	lp.zl.Panic().Msg(msg)
}

func (lp *zerologLogger) logEvent(e *zerolog.Event, props []log.Prop) {
	if len(props) == 0 {
		return
	}

	zle := &zlEvent{e: e}
	for _, p := range props {
		p(zle)
	}
}

func (lp *zerologLogger) Debug(msg string, props ...log.Prop) {
	if !lp.shouldLog(log.DebugLevel) {
		return
	}
	lp.logEvent(lp.zl.Debug(), props)
}

func (lp *zerologLogger) Info(msg string, props ...log.Prop) {
	if !lp.shouldLog(log.InfoLevel) {
		return
	}
	lp.logEvent(lp.zl.Info(), props)
}

func (lp *zerologLogger) Warn(msg string, props ...log.Prop) {
	if !lp.shouldLog(log.WarnLevel) {
		return
	}
	lp.logEvent(lp.zl.Warn(), props)
}

func (lp *zerologLogger) Error(msg string, props ...log.Prop) {
	if !lp.shouldLog(log.ErrorLevel) {
		return
	}
	lp.logEvent(lp.zl.Error(), props)
}

func (lp *zerologLogger) Panic(msg string, props ...log.Prop) {
	if !lp.shouldLog(log.PanicLevel) {
		return
	}
	lp.logEvent(lp.zl.Panic(), props)
}
