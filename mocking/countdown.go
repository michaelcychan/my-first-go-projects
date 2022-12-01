package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const startDeclaration string = "GO!"
const countDownCounter int = 3

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (ds *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countDownCounter; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, startDeclaration)
}

// this is a wrong implementation because of wrong sequence
// func Countdown(out io.Writer, sleeper Sleeper) {
// 	for i := countDownCounter; i > 0; i-- {
// 		fmt.Fprintln(out, i)
// 	}

// 	for i := countDownCounter; i > 0; i-- {
// 		sleeper.Sleep()
// 	}

// 	fmt.Fprint(out, startDeclaration)
// }

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
