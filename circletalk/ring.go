package main

import "sync"

type Element interface {
	Step()
	Val() int
}

type RingElement struct {
	val     int
	in, out chan int
}

func (e *RingElement) Step() {
	sent := make(chan struct{})
	go func() {
		defer close(sent)
		e.out <- e.val
	}()
	recv := <-e.in
	<-sent
	e.val = recv
}

func (e RingElement) Val() int {
	return e.val
}

type Ring struct {
	elements []Element
}

func (r Ring) Elements() []Element {
	return r.elements
}

func (r *Ring) Step() {
	var wg sync.WaitGroup
	wg.Add(len(r.Elements()))
	for _, el := range r.Elements() {
		go func(el Element) {
			defer wg.Done()
			el.Step()
		}(el)
	}
	wg.Wait()
}

func MakeRing(els int) Ring {
	elems := make([]Element, els)

	firstOut := make(chan int)
	lastOut := firstOut
	for i := 0; i < els; i++ {
		re := new(RingElement)
		re.val = i
		re.in = lastOut
		if i < els-1 {
			re.out = make(chan int)
		} else {
			re.out = firstOut
		}
		lastOut = re.out
		elems[i] = re
	}
	return Ring{elements: elems}
}
