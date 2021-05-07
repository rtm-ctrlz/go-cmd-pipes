package runner

import (
	"github.com/bketelsen/logr"
	"io"
	"io/ioutil"
	"os/exec"
	"sync"
)

func readPipe(wg *sync.WaitGroup, src *io.ReadCloser, dst *string) {
	wg.Add(1)
	var buf []byte
	buf, _ = ioutil.ReadAll(*src)
	*dst = string(buf)
	buf = make([]byte, 0, 0)
	wg.Done()
}

func RunIoUtilWGoRoutines(logger logr.Logger, cmd *exec.Cmd) (string, string, int) {
	logger.Info("Running RunIoUtilWGoRoutines")

	var wg sync.WaitGroup

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	var bStdout string
	var bStderr string
	var code = 0

	// start pipe-reading goroutines
	go readPipe(&wg, &stdout, &bStdout)
	go readPipe(&wg, &stderr, &bStderr)

	// start worker
	cmd.Start()

	// waiting for worker
	err := cmd.Wait()
	if err != nil {
		logger.Errorf("Command Error: \"%v\"", err)
		code = 1
	}

	// waiting for pipe-reading goroutines
	wg.Wait()

	return bStdout, bStderr, code
}
