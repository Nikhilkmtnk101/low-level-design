package main

import "fmt"

type Folder struct {
	name       string
	components []FileSystem
}

func NewFolder(name string, components []FileSystem) *Folder {
	return &Folder{
		name:       name,
		components: components,
	}
}

func (f *Folder) ls() {
	fmt.Printf("%s\n", f.name+":")
	for _, component := range f.components {
		component.ls()
	}
}
