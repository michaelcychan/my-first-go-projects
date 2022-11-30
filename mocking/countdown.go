package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const startDeclaration string = "GO!"
const countDownCounter int = 3

func Countdown(out io.Writer) {
	for i := countDownCounter; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(out, startDeclaration)
}

func main() {
	Countdown(os.Stdout)
}
