package rpc

import (
	"log"
	"os"
)

const (
	logPrefix = "chia/rpc: "
)

func init() {
	var (
		logErr = log.Default
		log    = log.New(os.StdOut, logPrefix, log.LstdFlags)
	)
	logErr.SetPrefix(logPrefix)
}
