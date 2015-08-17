package main

import "net/http"
import "bytes"
import "fmt"
import "io/ioutil"
import "encoding/json"
import "os"

func GetOozieServiceVersion() {
	resp, err := http.Get(fmt.Sprintf("%s/oozie/v1/admin/build-version", OozieURL()))
	check(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	var versionData map[string]interface{}
	json.Unmarshal(body, &versionData)
	log.Info(fmt.Sprintf("Connected to Oozie Version %s", versionData["buildVersion"]))
}

func PostNewWorkflow(wf ExecutableWorkflow) {
	log.Info(fmt.Sprintf("Submitting workflow %s", wf.Name))
	config := bytes.NewBuffer(wf.BuildConfiguration())
	resp, err := http.Post(fmt.Sprintf("%s/oozie/v1/jobs?action=start", OozieURL()), "application/xml", config)
	check(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	log.Info(string(body))
}

func OozieURL() string {
	return os.Getenv("OOZIE_URL")
}
