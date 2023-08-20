package node_tree

import (
	"fmt"
	"strconv"
	"strings"
)

type NodeType string
type CommandType string

const (
	Dir  NodeType = "dir"
	File          = "file"
)

const (
	Cd CommandType = "cd" // $ cd
	Ls             = "ls" // $ ls
)

type Command struct {
	Type       CommandType
	Args       string
	currentDir *Node
	parentDir  *Node
}

func ProcessCommand(line string, currentDir *Node, parentDir *Node) *Command {
	var command *Command
	if line[2] == 'c' && line[3] == 'd' {
		command = &Command{
			Type:       Cd,
			Args:       strings.Split(line, " ")[2],
			currentDir: currentDir,
			parentDir:  parentDir,
		}
	} else if line[2] == 'l' && line[3] == 's' {
		command = &Command{
			Type:       Ls,
			Args:       "",
			currentDir: currentDir,
			parentDir:  parentDir,
		}
	}
	return command
}

type Node struct {
	Name     string
	Type     NodeType
	Parent   *Node
	Children []*Node
	Size     int
	DirSize  int
}

func CreateFileNode(line string, parent *Node) *Node {
	sizeName := strings.Split(line, " ")
	size := sizeName[0]
	name := sizeName[1]
	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		fmt.Println("Error converting string to int")
		sizeInt = 0
	}
	return &Node{
		Name:     name,
		Type:     File,
		Parent:   parent,
		Children: nil,
		Size:     sizeInt,
		DirSize:  0,
	}
}

func CreateDirNode(dirName string, parent *Node) *Node {
	return &Node{
		Name:     dirName,
		Type:     Dir,
		Parent:   parent,
		Children: make([]*Node, 0),
		Size:     0,
		DirSize:  0,
	}
}

func (n *Node) GetSize() int {
	if n.Type == File {
		return n.Size
	} else {
		if n.DirSize == 0 {
			for _, child := range n.Children {
				n.DirSize += child.GetSize()
			}
		}
		return n.DirSize
	}
}

func ProcessTree(lines []string) {
	total := 0
	// create a hash map with node name and node
	nodeMap := make(map[string]*Node)
	var currentDir *Node
	lineNo := 0
	for _, line := range lines {
		if line[0] == '$' {
			// process command
			var parentDir *Node
			if currentDir != nil {
				parentDir = currentDir.Parent
			} else {
				parentDir = nil
			}
			command := ProcessCommand(line, currentDir, parentDir)
			if command.Type == Cd {
				if command.Args == "/" {
					root := nodeMap["root"]
					if root == nil {
						currentDir = &Node{
							Name:     "root",
							Type:     Dir,
							Parent:   nil,
							Children: make([]*Node, 0),
							Size:     0,
							DirSize:  0,
						}
						nodeMap["root"] = currentDir
					} else {
						currentDir = root
						parentDir = nil
					}
				} else {
					// check if the directory is already created
					parentDir = currentDir
					dir := nodeMap[command.Args]
					if dir == nil {
						// create a new directory
						dir = CreateDirNode(command.Args, parentDir)
						nodeMap[command.Args] = dir
					}
					currentDir = dir
				}
			} else if command.Type == Ls {
				if currentDir == nil {
					fmt.Println("currentDir is nil")
					continue
				}
				// iterate directory and append children
				for i := lineNo + 1; i < len(lines); i++ {
					contentLine := lines[i]
					if contentLine[0] == '$' {
						// go to process the next command
						break
					} else if contentLine[0] == 'd' &&
						contentLine[1] == 'i' &&
						contentLine[2] == 'r' {
						// directory
						directoryNode := CreateDirNode(contentLine[4:], currentDir)
						currentDir.Children = append(currentDir.Children, directoryNode)
					} else {
						// file
						fileNode := CreateFileNode(contentLine, currentDir)
						currentDir.Children = append(currentDir.Children, fileNode)
					}
				}
				if currentDir.GetSize() <= 100000 {
					fmt.Println("currentDir: ", currentDir.Name)
					fmt.Println("currentDir size: ", currentDir.GetSize())
					total += currentDir.GetSize()
				}
			}
		}
		lineNo++
	}
	fmt.Println("root dir: ", nodeMap["root"].Name)
	fmt.Println("number of children: ", len(nodeMap["root"].Children))
	fmt.Println("total size: ", total)
}
