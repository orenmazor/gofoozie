package main

import "os"
import "fmt"
import "io/ioutil"
import "encoding/xml"

type Coordinator struct {
	XMLName xml.Name "coordinator"
	Path    AppPath  `xml:"app-path"`
}

type AppPath struct {
	XMLName xml.Name "app-path"
}

type BundleApp struct {
	XMLName      xml.Name      "coordinator"
	Coordinators []Coordinator `xml:"coordinator"`
}

func loadValidJobs(files []os.FileInfo) {
	for _, element := range files {

		//read the xml
		xmlFile, err := os.Open(fmt.Sprintf("data/bundles/%s", element.Name()))
		check(err)
		defer xmlFile.Close()

		b, err := ioutil.ReadAll(xmlFile)
		check(err)

		var q BundleApp
		xml.Unmarshal(b, &q)
		log.Info(fmt.Sprintf("Found %s with %d coordinator jobs", element.Name(), len(q.Coordinators)))

	}
}

func DetectJobs() {
	bundlePath := "./data/bundles"
	// find all of the available bundles
	files, err := ioutil.ReadDir(bundlePath)
	check(err)
	// read all of the bundles and find all of the coordinator entries
	loadValidJobs(files)
	// read every coordinator entry and schedule a job for it with cron
}
