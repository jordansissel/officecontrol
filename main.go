package main

import (
	"time"

	"./denon"
)

func main() {
	denon.Discover()

	time.Sleep(10 * time.Second)
}
