package main

import (
	"flag"
	log "github.com/corvus-ch/logr/buffered"
	"go-cmd-pipes/memory"
	"go-cmd-pipes/runner"
	"go-cmd-pipes/worker"
	"os"
)

func main() {
	l := log.New(0)
	defer l.Buf().WriteTo(os.Stdout)

	logger := l.NewWithPrefix("[Manager] ")
	memLogger := logger.NewWithPrefix("[MEM] ")

	fSet := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	runSize := fSet.Int("s", 100, "Size of payload output, should be greater than 10")
	runAsWorker := fSet.Bool("w", false, "Run as worker")
	runType := fSet.String("t", "", "Type of handler: \"buf\" or \"io\" or \"io2\"")
	runDump := fSet.Bool("d", false, "Dump stderr/stdout of runner?")
	fSet.Parse(os.Args[1:])

	if len(*runType) == 0 && !*runAsWorker {
		fSet.Usage()
		return
	}

	if *runSize < 10 {
		logger.Error("Size (-s) should be greater than 10!")
		return
	}

	if *runAsWorker {
		worker.Run(*runSize)
		return
	}

	var bStdout string = ""
	var bStderr string = ""
	var code int = 0

	switch *runType {
	case "buf":
		bStdout, bStderr, code = runner.RunByteBuffers(logger, worker.GetCommand(*runSize))
		break
	case "io":
		bStdout, bStderr, code = runner.RunIoUtil(logger, worker.GetCommand(*runSize))
		break
	case "io2":
		bStdout, bStderr, code = runner.RunIoUtilWGoRoutines(logger, worker.GetCommand(*runSize))
		break
	default:
		logger.Errorf("Unsupported run type: \"%s\"", *runType)
		return
	}
	memory.PrintMemUsage(memLogger)

	logger.Infof("OutLen %d ErrLen %d Code %d", len(bStdout), len(bStderr), code)
	if *runDump {
		logger.Infof("StdOut: '%v'", bStdout)
		logger.Infof("StdErr: '%v'", bStderr)
	}
}
