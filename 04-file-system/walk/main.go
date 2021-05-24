package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type config struct {
	// extenstion to filter out
	ext string
	// min file size
	size int64
	// list files
	list bool
}

func main() {
	// Parsing command line flags
	root := flag.String("root", ".", "Root directory to start")
	// Action options
	list := flag.Bool("list", false, "List files only")
	// Filter options
	ext := flag.String("ext", "", "File extension to filter out")
	size := flag.Int64("size", 0, "Minimum files size")

	flag.Parse()

	c := config{
		ext:  *ext,
		size: *size,
		list: *list,
	}

	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
func run(root string, out io.Writer, cfg config) error {
	return nil
}
