package day7

type directory struct {
	name     string
	parent   *directory
	children map[string]*directory
	files    map[string]*file
}

func newDirectory(parent *directory, name string) *directory {
	dir := &directory{name: name, parent: parent}
	if dir.parent != nil {
		dir.parent.addDir(dir, name)
	}
	return dir
}

func (dir *directory) size() int {
	s := 0
	for _, f := range dir.files {
		s += f.size
	}
	for _, child := range dir.children {
		s += child.size()
	}
	return s
}

func (dir *directory) parentDir() *directory {
	return dir.parent
}

func (dir *directory) rootDir() *directory {
	d := dir
	for d.parent != nil {
		d = d.parent
	}
	return d
}

func (dir *directory) changeDir(name string) *directory {
	if target, ok := dir.children[name]; ok {
		return target
	}
	// Creates the directory if it doesn't already exist
	return newDirectory(dir, name)
}

func (dir *directory) addDir(child *directory, name string) {
	if dir.children == nil {
		dir.children = map[string]*directory{}
	} else if _, ok := dir.children[name]; ok {
		return
	}
	dir.children[name] = child
}

func (dir *directory) addFile(f *file) {
	if dir.files == nil {
		dir.files = map[string]*file{}
	} else if _, ok := dir.files[f.name]; ok {
		return
	}
	dir.files[f.name] = f
}

type file struct {
	name string
	size int
}

func newFile(name string, size int) *file {
	return &file{name: name, size: size}
}
