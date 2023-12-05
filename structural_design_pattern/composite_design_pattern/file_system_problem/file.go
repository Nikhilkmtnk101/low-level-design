package main

import "fmt"

type File struct {
	name string
}

func NewFile(name string) *File {
	return &File{
		name: name,
	}
}

func (f *File) ls() {
	fmt.Printf("%s\n", "-"+f.name)
}
