package main

func check(err error) {
	if err != nil {
		log.Error(err)
	}
}
