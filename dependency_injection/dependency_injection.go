package dependencyinjection

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, name string) {
	/*
		How to test the Printf?
		Para poder testear el printf, debemos usar la dependencia que usa el Printf
	*/
	fmt.Fprintf(writer, "Hello, %s", name)
}
