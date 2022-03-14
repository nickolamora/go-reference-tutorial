package main

import (
	"fmt"
	"os"
)

// Defer - means the statement will be deferred from executing at that particular moment. i.e. At the end of the method before returning
// a value

//Panic - panicking the application, like an exception. The panics needs to be recovered if possible if not it'll crash

func main() {
	f := createFile("defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating...")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("Writing...")
	fmt.Fprintln(f, "something")
}

func closeFile(f *os.File) {
	fmt.Println("Closing...")
	err := f.Close()

	if err != nil {
		fmt.Println(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
