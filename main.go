package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lRSErasureCode
#include "RSErasureCode.c"
#include <stdlib.h>  // Include for C.free
*/
import "C"
import (
    "os"
    "fmt"
    "unsafe"
    "strconv"
)

func main() {
    // get input from cli as shift value
    input := os.Args[1]

    // Convert Go strings to C strings
    fmt.Println("Input:", input)
    cInput := C.CString(input)
    defer C.free(unsafe.Pointer(cInput))

	C.init()
	C.init_dec()
}