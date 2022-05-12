package main

import (
	"example-restapi/src/router"
	"example-restapi/src/util"
)

func main() {
	util.PrintInfo("Run Application...")
	router.RunRouter()
	defer util.PrintInfo("Exit Application...")
}
