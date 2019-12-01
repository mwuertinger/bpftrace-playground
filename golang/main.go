package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	// start a lot of goroutines calling myfunc indefinitely
	for i := 0; i < 1; i++ {
		go func() {
			buf := make([]byte, 2<<24, 2<<24)
			for {
				myfunc(buf)
			}
		}()
	}

	// sleep forever
	for {
		time.Sleep(time.Hour)
	}
}

// myfunc does some nonsense expensive operations so that we have something
// to measure with bpftrace.
func myfunc(buf []byte) {
	_, err := rand.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sha256.Sum256(buf))
}
