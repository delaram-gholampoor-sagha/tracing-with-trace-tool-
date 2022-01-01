package main

import (
	"log"
	"os"
	"runtime/trace"
)

// tracing allows you to pass context through your system and evaluate where you are being held up , whether its by a third-party API call , a slow messaging quue , or an O(n2) function . tracing will help you where the bottleneck resides

// the go tracing tool hooks into the goroutine scheduler so that it can produce meaningful information about goroutines.


// install golang-misc package => in fedora run this command : dnf isntall golang-misc
// after running the project : run this command => tool trace trace.out 
func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	ch := make(chan string)
	go func() {
		ch <- "Hi Gophers"
	}()
	<-ch
	log.Printf("Trace Completed")
}
