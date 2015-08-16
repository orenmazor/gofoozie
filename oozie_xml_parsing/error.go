package oozie_xml_parsing

func check(err error) {
	if err != nil {
		log.Error(err)
	}
}
