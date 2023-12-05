package main

import "fmt"

type Folder struct {
	name       string
	components []FileSystem
}

func (f *Folder) ls() {
	fmt.Println(f.name)
	for _, component := range f.components {
		component.ls()
	}
}
