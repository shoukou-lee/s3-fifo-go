package main

import (
	"github.com/shoukou-lee/s3-fifo-go/internal"
)

type element struct {
	key   string
	value string
}

// main function is the entry point of the program
func main() {

	sequence := []element{
		{"lime", "a"},
		{"coral", "b"},
		{"lime", "c"},
		{"green", "d"},
		{"red", "e"},
		{"gray", "f"},
		{"gray", "g"},
		{"lime", "h"},
		{"black", "i"},
		{"lime", "j"},
		{"red", "k"},
		{"lime", "l"},
		{"pink", "m"},
		{"lime", "n"},
	}

	s3Fifo := internal.NewS3Fifo(20)
	for _, s := range sequence {
		s3Fifo.GetOrElsePut(s.key, s.value)
		s3Fifo.Log()
	}
}
