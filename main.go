package main

import (
	log "github.com/Sirupsen/logrus"
)

var (
	logger = log.New()
)

func main() {
	logger.Info("starting")
}
