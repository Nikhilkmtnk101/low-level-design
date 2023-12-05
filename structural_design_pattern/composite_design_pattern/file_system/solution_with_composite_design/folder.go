package solution_with_composite_design

import "fmt"

type Folder struct {
	name string
	components []FileSystem
}

func (f *Folder) ls()  {
	fmt.Println(f.name)
	for component := range f.components:

}

