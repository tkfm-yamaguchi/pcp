package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("[ERROR] Base file name should be specified.")
		os.Exit(1)
	}

	fname := args[0]

	if cur, err := filepath.Abs("."); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(filepath.Join(cur, fname))
	}
}
