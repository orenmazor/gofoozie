package main

import "encoding/xml"

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
	XMLName      xml.Name      "bundle"
	Name         string        `xml:"name,attr"`
	Coordinators []Coordinator `xml:"coordinator"`
}

type CoordinatorApp struct {
	XMLName   xml.Name "coordinator-app"
	Name      string   `xml:"name,attr"`
	start     string   `xml:"start,attr"`
	end       string   `xml:"end,attr"`
	frequency string   `xml:"frequency,attr"`
	timezone  string   `xml:"timezone,attr"`
	Action    Action   `xml:"action"`
}

type Action struct {
	XMLName  xml.Name "action"
	Workflow Workflow `xml:"workflow"`
}

type Workflow struct {
	XMLName xml.Name "workflow"
	Path    AppPath  `xml:"app-path"`
}

type WorkflowAction struct {
	to   string `xml:"to,attr"`
	name string `xml:"name,attr"`
}

type WorkflowApp struct {
	XMLName xml.Name         "workflow-app"
	Name    string           `xml:"name,attr"`
	Start   WorkflowAction   `xml:"start"`
	End     WorkflowAction   `xml:"end"`
	Kill    WorkflowAction   `xml:"kill"`
	Actions []WorkflowAction `xml:"action"`
}

type Configuration struct {
	XMLName    xml.Name   "configuration"
	Properties []Property `xml:"property"`
}

type Property struct {
	XMLName xml.Name "property"
	Name    Name     `xml:"name"`
	Value   Value    `xml:"value"`
}

type Name struct {
	XMLName xml.Name "name"
	Content string   `xml:",chardata"`
}

type Value struct {
	XMLName xml.Name "value"
	Content string   `xml:",chardata"`
}
