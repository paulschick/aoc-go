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
	}
}

func CreateDirNode(line string, parent *Node) *Node {
	return &Node{
		Name:     line[4:],
		Type:     Dir,
		Parent:   parent,
		Children: make([]*Node, 0),
		Size:     0,
	}
}

func ProcessTree(lines []string) {
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
			fmt.Println("command: ", command)
			if line[2] == 'c' && line[3] == 'd' {
				fmt.Println("cd command: ", line)
				if line[5] == '/' {
					fmt.Println("cd root")
					root := nodeMap["root"]
					if root == nil {
						fmt.Println("root not found")
						nodeMap["root"] = &Node{
							Name:     "root",
							Type:     Dir,
							Parent:   nil,
							Children: make([]*Node, 0),
							Size:     0,
						}
						currentDir = nodeMap["root"]
						fmt.Println("root: ", nodeMap["root"])
						//break
					} else {
						fmt.Println("root found")
						fmt.Println("root: ", root)
					}
				}
			} else if line[2] == 'l' && line[3] == 's' {
				if currentDir == nil {
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
						directoryNode := CreateDirNode(contentLine, currentDir)
						currentDir.Children = append(currentDir.Children, directoryNode)
					} else {
						// file
						fileNode := CreateFileNode(contentLine, currentDir)
						currentDir.Children = append(currentDir.Children, fileNode)
					}
				}
			}
		}
		lineNo++
	}
	fmt.Println("current dir: ", currentDir)
}
