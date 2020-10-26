package test

import (
	"design-pattern/prototype"
	"testing"
)

func TestPrototype(t *testing.T) {
	file1 := &prototype.File{Name: "File1"}
	file2 := &prototype.File{Name: "File2"}
	file3 := &prototype.File{Name: "File3"}
	folder1 := &prototype.Folder{
		Children: []prototype.Inode{file1},
		Name:     "Folder1",
	}
	folder2 := &prototype.Folder{
		Children: []prototype.Inode{folder1, file2, file3},
		Name:     "Folder2",
	}
	folder2.Print("  ")
	cloneFolder := folder2.Clone()
	cloneFolder.Print("  ")

}
