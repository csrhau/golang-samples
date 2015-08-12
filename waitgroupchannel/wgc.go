package main

import (
	"fmt"
	"sync"
)

func waitChannel(runners int) (*sync.WaitGroup, <-chan struct{}) {
	var wg sync.WaitGroup
	wg.Add(runners)
	ch := make(chan struct{})
	go func() {
		wg.Wait()
		defer close(ch)
	}()
	return &wg, ch
}

func main() {
	n, s, e, w := make(chan bool), make(chan bool), make(chan bool), make(chan bool)
	wg, done := waitChannel(4)
	go func() { n <- true }()
	go func() { s <- true }()
	go func() { e <- true }()
	go func() { w <- true }()
	func() {
		for {
			select {
			case <-n:
				fmt.Println("Received from North")
				wg.Done()
			case <-s:
				fmt.Println("Received from South")
				wg.Done()
			case <-e:
				fmt.Println("Received from East")
				wg.Done()
			case <-w:
				fmt.Println("Received from West")
				wg.Done()
			case <-done:
				return
			}
		}
	}()
}
