/*
 * File: main.go
 * Date: Thursday, 15th August 2024 9:50:32 PM
 * Author: Hy Nguyen
 * Copyright - 2024, HYSOFTWARE LLC
 */
package main

import (
	"os"

	"github.com/daiselabs/dl-go/pkg/dcfg"
	"github.com/daiselabs/dl-go/pkg/dlog"
	"github.com/daiselabs/dl-go/pkg/dlog/log"
)

func main() {
	dlog.SetRootLevel(log.DebugLevel)
	l := dlog.NewLogger(os.Stderr)
	l.InfoMsg("Simple Cfg Sample")
	dcfg.InitConfig(".")

	abc := dcfg.Cfg.GetString("ABC")
	l.InfoMsg("ABC = " + abc)

	f, err := os.Create("sample.log")
	if err != nil {
		l.Error("Failed to open file", dlog.AnErr("Error", err))
		return
	}

	defer f.Close()
	dlog.SetRootLevel(log.InfoLevel)
	fl := dlog.NewLogger(f)
	fl.DebugMsg("This is a debug message")
	fl.InfoMsg("There should be no DEBUG message!")

	cl := fl.With().Bool("FlagA", true).Logger()
	cl.InfoMsg("Message 1")
	cl.WarnMsg("Message 2")
}
