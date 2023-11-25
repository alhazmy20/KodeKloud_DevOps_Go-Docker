package main

import (
	"log"
	"sync"

	api "github.com/alhazmy20/Go-WebAPI/routes"
)

func main() {
	var wg sync.WaitGroup
	router := api.Initialize()
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := api.StartServer(router); err != nil {
			log.Fatal(err)
		}
	}()
	wg.Wait()
}
