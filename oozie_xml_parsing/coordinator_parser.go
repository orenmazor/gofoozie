package oozie_xml_parsing

import "os"
import "fmt"
import "io/ioutil"

func LoadCoordinator(bundleName string, coordinatorName string, path string) {
	xmlFile, err := os.Open(fmt.Sprintf("%s/coordinator.xml", path))
	check(err)
	defer xmlFile.Close()

	_, err = ioutil.ReadAll(xmlFile)
	check(err)

}
