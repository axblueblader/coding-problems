package leetcode

// https://leetcode.com/problems/design-in-memory-file-system/

import (
	"slices"
	"strings"
)

type FileSystem struct {
	root *FsNode
}

type FsNode struct {
	content  string
	children map[string]*FsNode
}

func Constructor() FileSystem {
	return FileSystem{
		root: &FsNode{
			children: make(map[string]*FsNode),
		},
	}
}

func (this *FileSystem) Ls(path string) []string {
	dirs := strings.Split(path, "/")

	current := this.root
	for _, dir := range dirs {
		dir = dir + "/"
		if current.children[dir] != nil {
			current = current.children[dir]
		} else {
			break
		}
	}

	if current.children[dirs[len(dirs)-1]] != nil {
		return []string{dirs[len(dirs)-1]}
	}

	answer := []string{}
	for k := range current.children {
		answer = append(answer, strings.Trim(k, "/"))
	}
	slices.Sort(answer)
	return answer
}

func (this *FileSystem) Mkdir(path string) {
	dirs := strings.Split(path, "/")

	current := this.root
	for _, dir := range dirs {
		dir = dir + "/"
		if current.children[dir] == nil {
			current.children[dir] = &FsNode{
				children: make(map[string]*FsNode),
			}
		}
		current = current.children[dir]
	}
}

func (this *FileSystem) AddContentToFile(filePath string, content string) {
	dirs := strings.Split(filePath, "/")
	fileName := dirs[len(dirs)-1]
	dirs = dirs[:len(dirs)-1]

	current := this.root
	for _, dir := range dirs {
		dir = dir + "/"
		if current.children[dir] == nil {
			current.children[dir] = &FsNode{
				children: make(map[string]*FsNode),
			}
		}
		current = current.children[dir]
	}

	if current.children[fileName] == nil {
		current.children[fileName] = &FsNode{
			content: content,
		}
	} else {
		current.children[fileName].content = current.children[fileName].content + content
	}
}

func (this *FileSystem) ReadContentFromFile(filePath string) string {
	dirs := strings.Split(filePath, "/")
	fileName := dirs[len(dirs)-1]
	dirs = dirs[:len(dirs)-1]

	current := this.root
	for _, dir := range dirs {
		dir = dir + "/"
		if current.children[dir] == nil {
			current.children[dir] = &FsNode{
				children: make(map[string]*FsNode),
			}
		}
		current = current.children[dir]
	}

	if current.children[fileName] != nil {
		return current.children[fileName].content
	}

	return ""
}

/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ls(path);
 * obj.Mkdir(path);
 * obj.AddContentToFile(filePath,content);
 * param_4 := obj.ReadContentFromFile(filePath);
 */
