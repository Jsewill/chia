package rpc

import (
	"log"
	"os"
)

const (
	logPrefix = "chia/rpc: "
)

var (
	logErr *log.Logger
)

func init() {
	var (
		logErr = log.Default()
		logger = log.New(os.Stdout, logPrefix, log.LstdFlags)
	)
	_ = logger
	logErr.SetPrefix(logPrefix)
}
