package main

import "encoding/xml"

type ExecutableWorkflow struct {
	Name string
	Path string
}

func (wf ExecutableWorkflow) Execute() {
	GetOozieServiceVersion()
	PostNewWorkflow(wf)
}

func (wf ExecutableWorkflow) BuildConfiguration() []byte {
	username := Property{"user.name", "oozie"}
	wfpath := Property{"oozie.wf.application.path", wf.Path}
	xmlbuffer, err := xml.Marshal(Configuration{Properties: []Property{username, wfpath}})
	check(err)

	return xmlbuffer
}

func (wf ExecutableWorkflow) Submit() {
}

func CreateExecutableWorkflow(name string, wfPath string) *ExecutableWorkflow {
	wf := new(ExecutableWorkflow)
	wf.Name = name
	wf.Path = wfPath
	return wf
}
