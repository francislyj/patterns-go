package composite

import "fmt"

type file struct {
	name string
}

func (f *file) search(content string)  {
	fmt.Printf("search content is %s in file %s\n", content, f.name)
}

func (f *file) getName() string {
	return f.name
}