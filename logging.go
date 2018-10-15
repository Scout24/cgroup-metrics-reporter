package main

import (
	"io/ioutil"
	"log"
	"os"
)

func verboseLogging(enable bool) {
	if !enable {
		log.SetOutput(ioutil.Discard)
	} else {
		log.SetOutput(os.Stdout)
	}
}
