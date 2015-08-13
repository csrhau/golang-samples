package main

import "sync"

type Element struct {
	val     int
	in, out chan int
}

func (e *Element) Step() {
	sent := make(chan struct{})
	go func() {
		defer close(sent)
		e.out <- e.val
	}()
	recv := <-e.in
	<-sent
	e.val = recv
}

type Ring struct {
	elements []*Element
}

func (r Ring) Elements() []*Element {
	return r.elements
}

func (r *Ring) Step() {
	var wg sync.WaitGroup
	wg.Add(len(r.Elements()))
	for _, el := range r.Elements() {
		go func(el *Element) {
			defer wg.Done()
			el.Step()
		}(el)
	}
	wg.Wait()
}

func MakeRing(els int) Ring {
	elems := make([]*Element, els)
	for i := 0; i < els; i++ {
		elems[i] = new(Element)
		elems[i].val = i
		elems[i].out = make(chan int)
		if i > 0 {
			elems[i].in = elems[i-1].out
		}
	}
	elems[0].in = elems[els-1].out
	return Ring{elements: elems}
}
