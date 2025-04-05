package main

import (
	"fmt"
	//"log"
	"KITSCDrafter/backend"
)

func main() {
	fmt.Println("Starting the application...")

	backend.SetUpDB()

}
