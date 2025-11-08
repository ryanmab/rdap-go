package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/ryanmab/rdap-go/pkg/client"
)

func main() {
	client := client.New()

	start := time.Now()
	response, err := client.LookupDomain("ryanmaber.co.uk")
	log.Printf("Lookup took: %s", time.Since(start))

	if err != nil {
		log.Panic(err)
	}

	data, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Panic(err)
	}

	log.Print(string(data))
}
