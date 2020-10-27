package main

import (
	"fmt"
	"portScanner/port"
)

func main() {
	fmt.Println("Port Scanning")
	results := port.InitialScan("172.16.4.13")
	fmt.Println(results)
}
