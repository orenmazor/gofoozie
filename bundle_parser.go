package main

import "os"
import "fmt"
import "io/ioutil"
import "encoding/xml"

func LoadBundles(files []os.FileInfo) []*ExecutableWorkflow {
	bundleJobs := make([]*ExecutableWorkflow, 0)
	for _, element := range files {

		//read the xml
		xmlFile, err := os.Open(fmt.Sprintf("data/bundles/%s", element.Name()))
		check(err)
		defer xmlFile.Close()

		b, err := ioutil.ReadAll(xmlFile)
		check(err)

		var q BundleApp
		xml.Unmarshal(b, &q)

		for _, coord := range q.Coordinators {
			bundleJobs = append(bundleJobs, LoadCoordinator(q.Name, coord.Name, coord.Path.Path))
		}

		log.Info(fmt.Sprintf("Found %s with %d jobs", q.Name, len(q.Coordinators)))

	}
	return bundleJobs
}
