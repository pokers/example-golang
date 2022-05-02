package main

import (
	pLog "example-logmodule/modules/log"
)
func main(){
	var index int = 1
	pLog.PrintInfo("[%d] Info level print... %s\n", index, "first")
	pLog.PrintDebug("[%d] Debug level print... %s\n", index, "second")
	pLog.PrintWarning("[%d] Warning level print... %s\n", index, "third")
	pLog.PrintError("[%d] Error level print... %s\n", index, "forth")
	pLog.PrintMsg("[%d] Message level print... %s\n", index, "fifth")
}
