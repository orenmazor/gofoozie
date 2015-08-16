package oozie_xml_parsing

import "github.com/Sirupsen/logrus"
import "os"
import "fmt"
import "io/ioutil"
import "encoding/xml"

var log = logrus.New()

type AppPath struct {
	XMLName xml.Name "app-path"
	Path    string   `xml:",chardata"`
}

type Coordinator struct {
	XMLName xml.Name "coordinator"
	Path    AppPath  `xml:"app-path"`
	Name    string   `xml:"name,attr"`
}

type BundleApp struct {
	XMLName      xml.Name      "coordinator"
	Name         string        `xml:"name,attr"`
	Coordinators []Coordinator `xml:"coordinator"`
}

func LoadBundles(files []os.FileInfo) {
	for _, element := range files {

		//read the xml
		xmlFile, err := os.Open(fmt.Sprintf("data/bundles/%s", element.Name()))
		check(err)
		defer xmlFile.Close()

		b, err := ioutil.ReadAll(xmlFile)
		check(err)

		var q BundleApp
		xml.Unmarshal(b, &q)
		log.Info(fmt.Sprintf("Found Bundle %s with %d coordinator jobs", q.Name, len(q.Coordinators)))

		for _, coord := range q.Coordinators {
			LoadCoordinator(q.Name, coord.Name, coord.Path.Path)
		}

	}
}
