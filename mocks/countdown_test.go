package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	/*
		Necesitaremos testear un iterative.
		De esta forma evitamos hackear y tendremos un código más limpio.
		Es importante dividir los problemas en pequeños para tener un software
		funcional
	*/
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := "3"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
