package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "world", "your name")
	flag.Parse()
	fmt.Printf("Hello, %s!\n", *name)
}
