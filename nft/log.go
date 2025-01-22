package nft

import (
	"log"
	"os"
)

const (
	logPrefix = "chia/nft: "
)

func init() {
	var (
		logErr = log.Default()
		logger = log.New(os.Stdout, logPrefix, log.LstdFlags)
	)
	logErr.SetPrefix(logPrefix)
}
