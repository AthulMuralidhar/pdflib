package main




// #cgo LDFLAGS: -L /home/athul/pdflib/pdflib10/lib/ -lpdf-linux
// #include "pdflib.h"
// #include "golang.h"
import "C"


func main() {
	_ = C.PDF_new();

}