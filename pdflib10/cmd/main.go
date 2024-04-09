package main

// #cgo LDFLAGS: -L /home/athul/pdflib/pdflib10/lib/ -lpdf-linux -lm  -lstdc++
// #include "pdflib.h"
// #include "golang.h"
import "C"

func main() {
	// -lm is for math.h
	_ = C.PDF_new();

}