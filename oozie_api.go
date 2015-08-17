package main

import "net/http"
import "fmt"
import "io/ioutil"
import "encoding/json"

func GetOozieServiceVersion() {
	resp, err := http.Get("http://hadoop-misc4.chi.shopify.com:11000/oozie/v1/admin/build-version")
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
	log.Info(wf.BuildConfiguration())
}
