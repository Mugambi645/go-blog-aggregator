// main.go
package main

import (
	"fmt"
	"log"

	" github.com/mugambi645/go-blog-aggregator" 
)

func main() {
	// 1. Read the initial config
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	// 2. Set the current user name and update the file on disk
	err = cfg.SetUser("patrick")
	if err != nil {
		log.Fatalf("error setting user: %v", err)
	}

	// 3. Read the config file again to verify changes
	updatedCfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading updated config: %v", err)
	}

	// 4. Print the structure contents to the terminal
	fmt.Printf("Database URL: %s\n", updatedCfg.DBURL)
	fmt.Printf("Current User: %s\n", updatedCfg.CurrentUserName)
}