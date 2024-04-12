package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	defer os.Exit(0)
	currentPath, err := os.Getwd()
	if err == nil {
		fmt.Printf("Current directory: %s\n", currentPath)
		fmt.Println("Input your age please. you confirmed the input by pressing enter.")
		var age int
		fmt.Scanln(&age)
		fmt.Println("正在分析中...")
		time.Sleep(time.Duration(2) * time.Second)
		fmt.Printf("Your age: %d\n", age)
		time.Sleep(time.Duration(2) * time.Second)
	} else {
		// handle error
		fmt.Println("error")
	}
}
