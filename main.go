package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lreed_solomon
#include "reed_solomon.c"
#include <stdlib.h>  // Include for C.free
*/
import "C"
import (
    "os"
    "fmt"
    "unsafe"
    // "strconv"
)

func main() {
    // get input from cli as message
    input := os.Args[1]

    // Convert Go strings to C strings
    fmt.Println("Original message:", input)
    cMessage := C.CString(input)
    defer C.free(unsafe.Pointer(cMessage))

    // Initialize the library
	C.init()
	C.init_dec()

    // set k
    C.set_k(512)

    // allocate memory for the output, the size of the output is 2^16*16 bits
    var cCodeword *C.char = (*C.char)(C.malloc(2<<16*16))
    defer C.free(unsafe.Pointer(cCodeword))

    // call the encode function
    C.encode(cMessage, cCodeword)
    fmt.Println("Message encoded.")

    // call the erasure_simulate function
    var cErasure *C._Bool = (*C._Bool)(C.malloc(2<<16*16))
    C.erasure_simulate(cCodeword, cErasure)
    fmt.Println("Erasure simulated.")

    // call the decode function
    C.decode(cCodeword, cErasure)
    
    // print the output
    fmt.Println("Decoded message:", C.GoString(cCodeword))
}