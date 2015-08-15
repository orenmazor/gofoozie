package main

import "github.com/Sirupsen/logrus"

var log = logrus.New()

func main() {
	log.Info("Hello. I'm GoFoozie. There is no Java here.")
	log.Info("Starting up the oozie runner...")
	StartOozie()
	log.Info("Starting up the web view....")
	RunWeb()
}
