package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const (
	finalWord      = "Go!"
	countdownStart = 3
	write          = "write"
	sleep          = "sleep"
)

type Sleeper interface {
	Sleep()
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	// Este approach es muy recomendado y bueno, permite muchas cosas y funciona e2e
	// El test demorará 4 segundos, lo que empeora la productividad
	for i := countdownStart; i > 0; i-- {
		// ¿La deferrencia del puntero se hace automática cuando se ejecuta un método?
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}
