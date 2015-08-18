package main

import "fmt"
import "io/ioutil"
import "time"

var jobs []*ExecutableWorkflow

func DetectJobs() {
	bundlePath := "./data/bundles"
	bundle_files, err := ioutil.ReadDir(bundlePath)
	check(err)

	jobs = LoadBundles(bundle_files)

	log.Info(fmt.Sprintf("Scheduled up %d jobs", len(jobs)))
}

func StartRunningJobs() {
	log.Info("Starting Job Runner")
	c := time.Tick(1 * time.Second)
	for range c {
		for _, job := range jobs {
			job.Execute()
		}
	}
}

func StartOozie() {
	DetectJobs()
	StartRunningJobs()
}
