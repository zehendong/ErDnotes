package main

import "os"
import "fmt"
import "simplemath"
import "strconv"

var Usage = func() {
	fmt.Println("USAVE: clac command [arguments] ...")
	fmt.Println("\nThe commands are:\n\tadd\tAddition of twovalues.\n\tsqrt\tSquare root aof a non-negative value.")
}

func main() {
	args := os.Args[1:]
	if args ==nil || len(args) < 2 {
		Usage()
		return
	}
}
