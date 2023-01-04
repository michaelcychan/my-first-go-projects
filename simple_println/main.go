package main

import (
	"fmt"
	"time"
)

func describe(x interface{}) {
	fmt.Printf("%%v is %v \n", x)
	fmt.Printf("%%#v is %#v \n", x)
	fmt.Printf("%%T is %T \n", x)
	fmt.Println("-------")
}

func main() {
	fmt.Println("Hello World! This is my first local Go Program!")
	fmt.Println("The time now is")
	fmt.Println(time.Now())
	fmt.Println("      `.-::::::-.`")
	fmt.Println("  .:-::::::::::::::-:.")
	fmt.Println("  `_:::    ::    :::_`")
	fmt.Println("   .:( ^   :: ^   ):.")
	fmt.Println("   `:::   (..)   :::.")
	fmt.Println("   `:::::::UU:::::::`")
	fmt.Println("   .::::::::::::::::.")
	fmt.Println("   O::::::::::::::::O")
	fmt.Println("   -::::::::::::::::-")
	fmt.Println("   `::::::::::::::::`")
	fmt.Println("    .::::::::::::::.")
	fmt.Println("      oO:::::::Oo")
	fmt.Print("This is Print")
	fmt.Print("This is second line of Print, they stay at same line")

	fmt.Println("")
	fmt.Println("-------------------")
	var int1 int = 42
	fmt.Println("information for int1")
	describe(int1)

	var str1 string = "This is a string"
	fmt.Println("information for str1")
	describe(str1)

	var bool1 bool = 1 == 2
	fmt.Println("information for bool1")
	describe(bool1)

	var bigint1 int64 = 420000
	fmt.Println("information for bigint1")
	describe(bigint1)
}
