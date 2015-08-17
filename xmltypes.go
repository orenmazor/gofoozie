package main

type AppPath struct {
	Path string `xml:",chardata"`
}

type Coordinator struct {
	Path AppPath `xml:"app-path"`
	Name string  `xml:"name,attr"`
}

type BundleApp struct {
	Name         string        `xml:"name,attr"`
	Coordinators []Coordinator `xml:"coordinator"`
}

type CoordinatorApp struct {
	Name      string `xml:"name,attr"`
	start     string `xml:"start,attr"`
	end       string `xml:"end,attr"`
	frequency string `xml:"frequency,attr"`
	timezone  string `xml:"timezone,attr"`
	Action    Action `xml:"action"`
}

type Action struct {
	Workflow Workflow `xml:"workflow"`
}

type Workflow struct {
	Path AppPath `xml:"app-path"`
}

type WorkflowAction struct {
	to   string `xml:"to,attr"`
	name string `xml:"name,attr"`
}

type WorkflowApp struct {
	Name    string           `xml:"name,attr"`
	Start   WorkflowAction   `xml:"start"`
	End     WorkflowAction   `xml:"end"`
	Kill    WorkflowAction   `xml:"kill"`
	Actions []WorkflowAction `xml:"action"`
}

type configuration struct {
	Properties []Property `xml:"property"`
}

type Property struct {
	Name  Name  `xml:"name"`
	Value Value `xml:"value"`
}

type Name struct {
	Content string `xml:",chardata"`
}

type Value struct {
	Content string `xml:",chardata"`
}
