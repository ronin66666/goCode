package main

import (
	"flag/cmd_utils/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err = %v", err)
	}
}
