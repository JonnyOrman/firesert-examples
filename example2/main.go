package main

import (
	"github.com/jonnyorman/firesert"
)

type FirestoreDocument struct {
	Prop1 string
	Prop2 int
}

func main() {
	firesert.RunTyped[FirestoreDocument]()
}
