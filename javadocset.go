package main

import (
	"flag"
	"fmt"
	"github.com/samcday/go-dash-javadocset"
	"os"
)

func main() {
	flag.Parse()
	var args = flag.Args()

	if len(args) != 2 {
		fmt.Println("Usage: go-dash DocSetName javadocDir")
		os.Exit(1)
	}
	fmt.Println(args)
	if err := javadocset.Build( args[1], ".", args[0]); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
