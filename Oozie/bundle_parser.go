package oozie

import "github.com/Sirupsen/logrus"
import "os"
import "fmt"
import "io/ioutil"
import "encoding/xml"

var log = logrus.New()

type Job struct {
	command string
}

func LoadBundles(files []os.FileInfo) []Job {
	for _, element := range files {

		//read the xml
		xmlFile, err := os.Open(fmt.Sprintf("data/bundles/%s", element.Name()))
		check(err)
		defer xmlFile.Close()

		b, err := ioutil.ReadAll(xmlFile)
		check(err)

		var q BundleApp
		xml.Unmarshal(b, &q)
		/* log.Info(fmt.Sprintf("Found Bundle %s with %d coordinator jobs", q.Name, len(q.Coordinators))) */

		for _, coord := range q.Coordinators {
			LoadCoordinator(q.Name, coord.Name, coord.Path.Path)
		}

	}
	return make([]Job, 4, 4)
}
