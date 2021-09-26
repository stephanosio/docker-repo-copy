package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 2 {
		fmt.Printf("usage: %s <src> <dest> [<dest>]\n", os.Args[0])
		os.Exit(1)
	}

	if err := copyAll(args[0], args[1:]); err != nil {
		panic(err)
	}
}

func copyAll(src string, dest []string) error {
	for _, d := range dest {
		if err := craneCopy(src, d); err != nil {
			return err
		}
	}
	return nil
}

func craneCopy(src string, dest string) error {
	return run([]string{"copy", src, dest})
}

func run(args []string) error {
	fmt.Println("crane", args)
	cmd := exec.Command("crane", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
