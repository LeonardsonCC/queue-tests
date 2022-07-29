package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/LeonardsonCC/queue-tests/crawler"
	"github.com/LeonardsonCC/queue-tests/server"
)

func main() {
	log.Println("Starting...")
	// starting a http server, bc it's cool
	srv := server.NewServer()
	go srv.Start()

	log.Println("Starting crawlers...")

	// how many appends and pops to run (8000 already makes my machine cry and show how incable david is)
	count := 8000

	// i know, wg is useless at the moment, but i was using ok
	var wg sync.WaitGroup

	// jorge adds everything he can
	wg.Add(1)
	jorge := crawler.NewCrawler("http://localhost:8080/queue/add", count)
	jorge.Run(&wg)

	// david tries to remove everything that jorge added, but he have just 1 second to do this
	wg.Add(1)
	david := crawler.NewCrawler("http://localhost:8080/queue/pop", count)
	go david.Run(&wg)

	time.Sleep(1 * time.Second)
	// wg.Wait()

	// how much of what jorge added still there, bc david was incable to remove all
	resp, _ := http.Get("http://localhost:8080/queue/length")
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("Length: ", string(body))
}
