package util

import (
	"encoding/json"
	"fmt"
)

func printLog(level int, msg string, values ...any) {
	var strlevel string = ""
	switch level {
	case 0:
		strlevel = "INFO"
	case 1:
		strlevel = "DEBUG"
	case 2:
		strlevel = "WARNING"
	case 3:
		strlevel = "ERROR"
	default:
		strlevel = "LOG"
	}

	fmt.Printf(fmt.Sprintf("%8s > ", strlevel)+msg, values...)
	fmt.Printf("\n")
}

func PrintInfo(msg string, values ...any) {
	printLog(0, msg, values...)
}
func PrintDebug(msg string, values ...any) {
	printLog(1, msg, values...)
}
func PrintWarning(msg string, values ...any) {
	printLog(2, msg, values...)
}
func PrintError(msg string, values ...any) {
	printLog(3, msg, values...)
}
func PrintMsg(msg string, values ...any) {
	printLog(4, msg, values...)
}
func PrintJson(title string, jsonObj any) {
	result, err := json.Marshal(jsonObj)
	if err == nil {
		fmt.Println(title, string(result))
	}
}
