package main

import "fmt"

type File struct {
	Name string
}

func (f *File) ls() {
	fmt.Println("- ", f.Name)
}
