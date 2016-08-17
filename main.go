package main

import (
	"encoding/json"
	"fmt"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	log "github.com/Sirupsen/logrus"
)

var (
	logger = log.New()
	config *configRoot
)

var (
	flagConfigPath     = kingpin.Flag("config", "Config file path.").Default("config.json").Short('c').File()
	flagDebug          = kingpin.Flag("debug", "Debug enabled.").Short('d').Bool()
	flagConfigTemplate = kingpin.Flag("template", "Print a template configuration file and exit.").Bool()
	//verbose = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	//name    = kingpin.Arg("name", "Name of user.").Required().String()
)

func main() {
	kingpin.Parse()
	if *flagConfigTemplate {
		c, _ := json.MarshalIndent(configDefault, "", "    ")
		fmt.Println(string(c))
		return
	}
	var err error
	config, err = configParseFile(*flagConfigPath)
	if err != nil {
		logger.Fatal("Unable to load config, ", err)
	}

	logger.Info("Webviewer starting")

}
