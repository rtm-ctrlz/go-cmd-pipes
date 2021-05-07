package worker

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetCommand(size int) *exec.Cmd {
	return exec.Command(os.Args[0], "-w", "-s", strconv.FormatInt(int64(size), 10))
}

func Run(size int) {
	var batchSize int = 50
	var batchCount int = int(math.Floor(float64(size) / float64(batchSize)))
	var batchRest int = size - batchSize*batchCount
	var written int = 0
	var strFmt = fmt.Sprintf("%%.%dd\n", batchSize-3)
	var str string
	for i := 0; i < batchCount; i++ {
		str = fmt.Sprintf(strFmt, written)
		fmt.Fprint(os.Stdout, "O:"+str)
		fmt.Fprint(os.Stderr, "E:"+str)
		written += batchSize
	}
	if batchRest > 0 {
		str = strings.Repeat("R", batchRest-3) + "\n"
		fmt.Fprint(os.Stdout, "O:"+str)
		fmt.Fprint(os.Stderr, "E:"+str)
	}
}
