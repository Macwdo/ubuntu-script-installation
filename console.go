package main

import (
	"fmt"
	"time"
)

func LogMessage(message string, style Style, color Color) {
	fmt.Printf("%s%s%s%s\n", style, color, message, Reset)
}

func LogMessagef(message string, style Style, color Color) {
	fmt.Printf("%s%s%s%s", style, color, message, Reset)
}

func Separator() {
	fmt.Println("---------------------------------------------------")
}

func LogMessageWithLoading(message string, style Style, color Color, loadingTime int) {
	if loadingTime == 0 {
		loadingTime = 3
	}
	fmt.Printf("%s%s%s%s", style, color, message, Reset)
	for i := 0; i < loadingTime; i++ {
		time.Sleep(time.Second * 1)
		fmt.Print(".")
	}
	fmt.Println()
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func WaitForKeyPress() {
	fmt.Scanln()
}
