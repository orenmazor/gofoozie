package main

import "github.com/Sirupsen/logrus"

var log = logrus.New()

func main() {
	log.Info("Hello. I'm GoFoozie. There is no Java here.")
	log.Info("Starting up the web view....")
	go RunWeb()
	log.Info("Starting up the oozie runner...")
	StartOozie()
}
