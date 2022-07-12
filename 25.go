package main

import (
	"log"
	"time"
)

func main() {
	log.Println("before sleep")
	sleep(time.Second * 2)
	log.Println("after sleep")
}

func sleep(t time.Duration) {
	<-time.After(t)
}
