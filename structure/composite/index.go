package composite

func TestComposite() {
	f1 := &file{
		name: "aa",
	}

	f2 := &file{
		name: "bb",
	}

	fs := folder{
		name: "ff",
	}

	fs.add(f1)
	fs.add(f2)

	fs.search("hello world")
}
