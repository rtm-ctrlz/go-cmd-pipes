package runner

import (
	"github.com/bketelsen/logr"
	"io/ioutil"
	"os/exec"
)

func RunIoUtil(logger logr.Logger, cmd *exec.Cmd) (string, string, int) {
	logger.Info("Running with ioutil")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		logger.Errorf("Get StdoutPipe error: \"%v\"", err)
		return "", "", 1
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		logger.Errorf("Get StderrPipe error: \"%v\"", err)
		return "", "", 1
	}

	if err = cmd.Start(); err != nil {
		logger.Errorf("Failed to start command: \"%v\"", err)
	}

	bStderr, _ := ioutil.ReadAll(stderr)
	bStdout, _ := ioutil.ReadAll(stdout)

	var code int = 0
	err = cmd.Wait()
	if err != nil {
		logger.Errorf("Command Error: \"%v\"", err)
		code = 1
	}

	return string(bStdout), string(bStderr), code
}
