package main

import "os"
import "fmt"
import "io/ioutil"
import "encoding/xml"

func LoadCoordinator(bundleName string, coordinatorName string, path string) *ExecutableWorkflow {
	xmlFile, err := os.Open(fmt.Sprintf("%s/coordinator.xml", path))
	check(err)
	defer xmlFile.Close()

	file, err := ioutil.ReadAll(xmlFile)
	check(err)

	var q CoordinatorApp
	xml.Unmarshal(file, &q)

	return LoadWorkflow(q.Name, q.Action.Workflow.Path.Path)
}
