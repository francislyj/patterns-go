package prototype

import "fmt"

type folder struct {
	children []inode
	name string
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)

	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *folder) clone() inode {
	cloneFolder := &folder{
		name: f.name,
	}

	var tempChildren []inode

	for _,i := range f.children {
		child := i.clone()
		tempChildren = append(tempChildren, child)
	}

	cloneFolder.children = tempChildren

	return cloneFolder

}