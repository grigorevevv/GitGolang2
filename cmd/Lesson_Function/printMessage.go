package main

import "fmt"

func printMessage(level int) {

	switch level {
	case 1:
		fmt.Println("Info message")
	case 2:
		fmt.Println("Warning message")
	case 3:
		fmt.Println("Error message")
	default:
		fmt.Println("Unknown level")
	}
}
