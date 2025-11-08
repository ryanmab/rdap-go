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
	response, err := client.LookupIPv4("8.8.8.8")
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
