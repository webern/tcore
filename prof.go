// Copyright (c) 2019 by Matthew James Briggs, https://github.com/webern

package tcore

import (
	"os"
	"runtime/pprof"
	"time"

	"github.com/bitflip-software/iqbid/blaze/flog"
)

func StartPPROF(outFilepath string) *os.File {
	if len(outFilepath) >= 255 {
		flog.Warn("filename too long")
		outFilepath = outFilepath[:255]
	}

	traceFile, err := os.Create(outFilepath)
	if err != nil {
		flog.Errorf(err.Error())
		os.Exit(1)
	} else if traceFile == nil {
		flog.Errorf("tracefile is nil")
		os.Exit(1)
	}

	err = pprof.StartCPUProfile(traceFile)
	if err != nil {
		flog.Errorf(err.Error())
		os.Exit(1)
	}

	return traceFile
}

func StopPPROF(f *os.File) {
	time.Sleep(10 * time.Millisecond)
	pprof.StopCPUProfile()
	time.Sleep(10 * time.Millisecond)
	err := f.Close()
	if err != nil {
		flog.Error(err.Error())
		os.Exit(1)
	}
}
