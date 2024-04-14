package main

import (
	"fmt"
	"time"
)

func script() {
	fmt.Println("Running the script...")
	time.Sleep(5 * time.Second)
	fmt.Println("Hello, World!")
}
