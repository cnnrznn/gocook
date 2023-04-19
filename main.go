package main

import "runtime"

func main() {
	for i := 0; i < runtime.NumCPU()-1; i++ {
		go cook()
	}
	cook()
}

func cook() {
	a := 0
	b := 1
	for {
		c := a + b
		a = b
		b = c
	}
}
