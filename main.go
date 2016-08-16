package main

import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

var (
	logger = log.New()
)

func main() {
	logger.Info("starting")
	c, _ := json.MarshalIndent(configDefault, "", "    ")
	fmt.Println(string(c))
}
