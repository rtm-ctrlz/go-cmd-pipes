package memory

import (
	"github.com/bketelsen/logr"
	"runtime"
)

func PrintMemUsage(logger logr.Logger) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	memLogger := logger.NewWithPrefix("[MEM] ")
	memLogger.Infof(
		"Alloc = %v MiB | TotalAlloc = %v MiB | Sys = %v MiB | NumGC = %v",
		bToMb(m.Alloc),
		bToMb(m.TotalAlloc),
		bToMb(m.Sys),
		m.NumGC,
	)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
