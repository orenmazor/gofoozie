package oozie

import "encoding/xml"

type ExecutableWorkflow struct {
	Name string
	Path string
}

func (wf ExecutableWorkflow) Execute() {
	GetOozieServiceVersion()
	PostNewWorkflow(wf)
}

func (wf ExecutableWorkflow) BuildConfiguration() string {
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
	config := Configuration{Properties: properties}

	xmlstr, err := xml.Marshal(config)
	check(err)

	return string(xmlstr)
}

func (wf ExecutableWorkflow) Submit() {
}

func CreateExecutableWorkflow(name string, xmlDoc WorkflowApp) *ExecutableWorkflow {
	wf := new(ExecutableWorkflow)
	wf.Name = name
	wf.Path = "asdf"
	return wf
}
