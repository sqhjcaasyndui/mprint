package Mprint

import (
	"fmt"

	"github.com/fatih/color"
)

func printInfo(message string) {

	info := color.New(color.FgGreen).SprintfFunc()
	tagStr := fmt.Sprintf("[%s]", info("info"))
	fmt.Printf("%-12s %s\n", tagStr, message)
}

func printErr(message string) {
	err := color.New(color.FgRed).SprintfFunc()
	tagStr := fmt.Sprintf("[%s]", err("err"))
	fmt.Printf("%-12s %s\n", tagStr, message)
}
