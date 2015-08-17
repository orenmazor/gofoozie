package main

import "os"
import "fmt"
import "io/ioutil"
import "encoding/xml"

func LoadWorkflow(name string, path string) *ExecutableWorkflow {
	xmlFile, err := os.Open(fmt.Sprintf("%s/workflow.xml", path))
	check(err)
	defer xmlFile.Close()

	file, err := ioutil.ReadAll(xmlFile)
	check(err)

	var q WorkflowApp
	xml.Unmarshal(file, &q)

	wf := CreateExecutableWorkflow(name, q)
	return wf
}
