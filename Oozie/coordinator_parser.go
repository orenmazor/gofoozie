package oozie

import "os"
import "fmt"
import "io/ioutil"
import "encoding/xml"

func LoadCoordinator(bundleName string, coordinatorName string, path string) []Workflow {
	xmlFile, err := os.Open(fmt.Sprintf("%s/coordinator.xml", path))
	check(err)
	defer xmlFile.Close()

	file, err := ioutil.ReadAll(xmlFile)
	check(err)

	var q CoordinatorApp
	xml.Unmarshal(file, &q)
	log.Info(fmt.Sprintf("Found %s::%s with %d workflows", bundleName, q.Name, len(q.Action.Workflows)))

	return make([]Workflow, 3, 3)
}
