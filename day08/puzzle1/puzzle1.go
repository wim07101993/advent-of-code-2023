package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Nodes map[string]Node
type Node struct {
	Left  string
	Right string
}

func main() {
	fmt.Println(Solve(os.Stdin))
}

func Solve(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	// scan the instructions
	scanner.Scan()
	instr := scanner.Text()

	// remove the white line in between
	scanner.Scan()

	nodes := Nodes{}
	for scanner.Scan() {
		name, node := ParseNode(scanner.Text())
		nodes[name] = node
	}

	return nodes.CalcSteps(instr)
}

func ParseNode(s string) (name string, node Node) {
	ss := strings.Split(s, " = ")
	name = ss[0]
	ns := strings.Split(strings.Trim(ss[1], "()"), ", ")
	return name, Node{ns[0], ns[1]}
}

func (ns Nodes) CalcSteps(instr string) int {
	name := "AAA"
	node, ok := ns[name]
	if !ok {
		panic(fmt.Errorf("did not find node AAA"))
	}

	sum := 0
	for {
		for i, lr := range instr {
			switch lr {
			case 'L':
				name = node.Left
			case 'R':
				name = node.Right
			}

			if name == "ZZZ" {
				return sum + i + 1
			}

			node, ok = ns[name]
			if !ok {
				panic(fmt.Errorf("did not find node AAA"))
			}
		}
		sum += len(instr)
	}
}
