package main

import "fmt"
import "io/ioutil"
import "./oozie_xml_parsing"

func DetectJobs() {
	bundlePath := "./data/bundles"
	// find all of the available bundles
	files, err := ioutil.ReadDir(bundlePath)
	check(err)
	// read all of the bundles and find all of the coordinator entries
	oozie_xml_parsing.LoadBundles(files)
	// read every coordinator entry and schedule a job for it with cron
}
func AddOozieCoordinator(bundleName string, coordinatorName string, appPath string) {
	log.Info(fmt.Sprintf("Adding another job for %s", appPath))
	/* log.Info("hello") */
}

func StartOozie() {
	DetectJobs()
}
