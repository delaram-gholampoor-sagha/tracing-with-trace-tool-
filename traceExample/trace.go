package main

import (
	"log"
	"os"
	"runtime/trace"
)




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
