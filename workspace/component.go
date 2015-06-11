package main

import (
	"fmt"
	"flag"
	"os"
)

var componentDir = "../static/component/"

var mode = flag.String("m", "add", "The mode to run in. Options are [add, list, list-all].")
var name = flag.String("n", "", "The name of the component to operate on.")
var version = flag.Int("v", -1, "")
var force = flag.Bool("f", false, "Whether to force overwriting if the version is specified.")

func main() {
	flag.Parse()

	if *mode == "add" {
		err := Add(*name, *version, *force)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func Add(name string, version int, force bool) error {
	if len(name) == 0 {
		return fmt.Errorf("Name (-n) must be set.")
	}
	componentCSS, err := os.OpenFile(name + ".css", os.O_RDONLY, 0600)
	if err != nil {
		return err
	}
	componentJS, err := os.OpenFile(name + ".js", os.O_RDONLY, 0600)
	if err != nil {
		return err
	}

	fmt.Println(componentCSS, componentJS)

	if version > 0 {
		// ?
	}

	return nil
}
