package oozie

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
	XMLName   xml.Name   "action"
	Workflows []Workflow `xml:"workflow"`
}

type Workflow struct {
	XMLName xml.Name "workflow"
	Path    AppPath  `xml:"app-path"`
}
