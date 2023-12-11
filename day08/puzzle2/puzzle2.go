package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
	"sync"
	"time"
)

type Node struct {
	Name  string
	Left  *Node
	Right *Node
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

	allNodes := ParseNodes(scanner)
	starters := GetStarters(allNodes)

	return CalcStepsParallel(starters, instr)
}

func ParseNodes(scanner *bufio.Scanner) []*Node {
	var ret []*Node
	nodes := map[string]*Node{}

	for scanner.Scan() {
		ss := strings.Split(scanner.Text(), " = ")
		ns := strings.Split(strings.Trim(ss[1], "()"), ", ")

		name := ss[0]
		leftName := ns[0]
		rightName := ns[1]

		node, ok := nodes[name]
		if !ok {
			node = &Node{Name: name}
			nodes[name] = node
			ret = append(ret, node)
		}

		left, ok := nodes[leftName]
		if !ok {
			left = &Node{Name: leftName}
			nodes[leftName] = left
			ret = append(ret, left)
		}

		right, ok := nodes[rightName]
		if !ok {
			right = &Node{Name: rightName}
			nodes[rightName] = right
			ret = append(ret, right)
		}

		node.Left = left
		node.Right = right

	}

	return ret
}

func GetStarters(ns []*Node) []*Node {
	var starters []*Node
	for _, n := range ns {
		if strings.HasSuffix(n.Name, "A") {
			starters = append(starters, n)
		}
	}
	return starters
}

func CalcStepsParallel(ns []*Node, instr string) int {
	stepCounts := make([]int, len(ns))
	for {
		done := false
		go func() {
			for !done {
				time.Sleep(time.Second)
				fmt.Println(stepCounts)
			}
		}()

		wg := sync.WaitGroup{}
		wg.Add(len(ns))
		for i := range ns {
			go func(i int) {
				for stepCounts[i] == 0 || stepCounts[i] < slices.Max(stepCounts) {
					fmt.Printf("counting %v, max: %v\n", stepCounts[i], slices.Max(stepCounts))
					steps, newNode := ns[i].CalcStepsUntilNextZ(instr)
					ns[i] = newNode
					stepCounts[i] += steps
				}
				wg.Done()
			}(i)
		}

		wg.Wait()
		done = true
		allTheSame := true
		for _, steps := range stepCounts {
			if steps != stepCounts[0] || steps == 0 {
				allTheSame = false
				break
			}
		}

		if allTheSame {
			return stepCounts[0]
		}
	}
}

func CalcSteps(ns []*Node, instr string) int {
	sum := 0
	for {
		for i, lr := range instr {
			allEndWithZ := true
			for i, node := range ns {
				var newNode *Node
				switch lr {
				case 'L':
					newNode = node.Left
				case 'R':
					newNode = node.Right
				}

				ns[i] = newNode
				if !strings.HasSuffix(newNode.Name, "Z") {
					allEndWithZ = false
				}
			}

			if allEndWithZ {
				return sum + i + 1
			}
		}
		sum += len(instr)
	}
}

func (n *Node) CalcStepsUntilNextZ(instr string) (int, *Node) {
	iterations := 0
	var newNode = n
	for {
		for i, lr := range instr {
			switch lr {
			case 'L':
				newNode = newNode.Left
			case 'R':
				newNode = newNode.Right
			}

			if strings.HasSuffix(newNode.Name, "Z") {
				return iterations + i + 1, newNode
			}
		}
		iterations += len(instr)
	}
}

func (n *Node) String() string {
	return "{Node: " + n.Name + " L: " + n.Left.Name + " R: " + n.Right.Name + "}"
}
