// Copyright (c) 2019 by Matthew James Briggs, https://github.com/webern

package tcore

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func StartPPROF(outFilepath string) *os.File {
	if len(outFilepath) >= 255 {
		fmt.Print("filename too long")
		outFilepath = outFilepath[:255]
	}

	traceFile, err := os.Create(outFilepath)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	} else if traceFile == nil {
		fmt.Printf("tracefile is nil")
		os.Exit(1)
	}

	err = pprof.StartCPUProfile(traceFile)
	if err != nil {
		fmt.Printf(err.Error())
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
		fmt.Printf(err.Error())
		os.Exit(1)
	}
}
