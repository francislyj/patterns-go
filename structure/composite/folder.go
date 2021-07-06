package composite

import "fmt"

type folder struct {
	components []component
	name string
}

func (f *folder) search(content string) {
	fmt.Printf("search data in folder content %s, folder name is %s\n", content, f.name)
	for _, com := range f.components {
		com.search(content)
	}
}

func (f *folder) add(com component) {
	f.components = append(f.components, com)
}
