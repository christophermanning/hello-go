package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	root, err := os.OpenRoot(path.Join(dir, "example"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// can access files under the isolated root directory
	_, err = root.Open("README.md")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dockerfile := path.Join(dir, "Dockerfile")

	// cannot access dockerfile using root.Open isolated to a different directory
	_, err = root.Open(dockerfile)
	if err != nil {
		fmt.Println("Cannot access file outside root path, as expected")
	}
	defer root.Close()

	// can access dockerfile using os.Open
	fs, err := os.Open(dockerfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fs.Close()
}
