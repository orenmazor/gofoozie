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
	name1 := "user.name"
	name2 := "oozie.wf.application.path"
	value1 := "oozie"
	value2 := wf.Path
	properties := make([]Property, 2)
	name := Name{Content: name1}
	value := Value{Content: value1}
	properties[0] = Property{Name: name, Value: value}
	name = Name{Content: name2}
	value = Value{Content: value2}
	properties[1] = Property{Name: name, Value: value}
	config := configuration{Properties: properties}

	xmlbuffer, err := xml.Marshal(config)
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
