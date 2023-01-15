package main

import (
	"io/ioutil"
	"os"
)

func main() {
	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	Trace.Println("I have something standard to say")
	Info.Println("Special Information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
}
