package runner

import (
	"bytes"
	"github.com/bketelsen/logr"
	"os/exec"
	"runtime"
)

func RunByteBuffers(logger logr.Logger, cmd *exec.Cmd) (string, string, int) {
	logger.Info("Running with RunByteBuffers")
	defer runtime.GC()
	var bErr bytes.Buffer
	var bOut bytes.Buffer
	cmd.Stdout = &bOut
	cmd.Stderr = &bErr

	err := cmd.Run()

	bStdout := bOut.Bytes()
	bStderr := bErr.Bytes()

	var code = 0
	if err != nil {
		logger.Errorf("Command Error: \"%v\"", err)
		code = 1
	}
	return string(bStdout), string(bStderr), code
}
