package main

import "fmt"
import "io/ioutil"
import "./Oozie"

func DetectJobs() {
	bundlePath := "./data/bundles"
	bundle_files, err := ioutil.ReadDir(bundlePath)
	check(err)

	jobs := oozie.LoadBundles(bundle_files)

	log.Info(fmt.Sprintf("Scheduled up %d jobs", len(jobs)))
}
func AddOozieCoordinator(bundleName string, coordinatorName string, appPath string) {
	log.Info(fmt.Sprintf("Adding another job for %s", appPath))
	/* log.Info("hello") */
}

func StartOozie() {
	DetectJobs()
}
