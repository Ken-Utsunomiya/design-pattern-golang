package main

import (
	"github.com/projects/design-pattern-golang/Structure/Composite/component"
	"github.com/projects/design-pattern-golang/Structure/Composite/composite"
)

func main() {
	file1 := &component.File{Name: "File1"}
	file2 := &component.File{Name: "File2"}
	file3 := &component.File{Name: "File3"}

	folder1 := &composite.Folder{Name: "Folder1"}

	folder1.Add(file1)

	folder2 := &composite.Folder{Name: "Folder2"}
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)

	folder2.Search("rose")
}
