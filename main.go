package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	name   int
	parent int
	left   int
	right  int
}

type Tree []node

func (tree *Tree) out(id int, writer *bufio.Writer) { // output function
	if (*tree)[id-1].left != 0 {
		tree.out((*tree)[id-1].left, writer)
	}

	fmt.Fprintf(writer, "%d ", (*tree)[id-1].name)

	if (*tree)[id-1].right != 0 {
		tree.out((*tree)[id-1].right, writer)
	}
}

const maxcap = 1024 * 1024 * 10

// реализация бинарного дерева, а также манипуляций с ним
// на основе входящих данных
func main() {
	fi, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("failed to open input.txt: %v", err)
		return
	}

	defer fi.Close()

	scanner := bufio.NewScanner(fi)
	buff := make([]byte, maxcap)
	scanner.Buffer(buff, maxcap)
	scanner.Scan()

	dataslice := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(dataslice[0]) //amount of nodes in the tree
	q, _ := strconv.Atoi(dataslice[1]) //amount of tree permutations

	scanner.Scan()
	data := scanner.Text() //getting permutations
	tree := make(Tree, n)
	dataslice = strings.Split(data, " ") // permutations string -> []string

	if len(dataslice) > q {
		dataslice = dataslice[:len(dataslice)-1]
	}

	mutations := make([]int, q)

	for d, i := range dataslice {
		mutations[d], _ = strconv.Atoi(i) // permutations []string -> []int
	}

	for i := 1; i <= n; i++ { // creating a complete tree
		switch {
		case 2*i+1 <= n:
			tree[i-1] = node{name: i, parent: i / 2, left: i * 2, right: i*2 + 1}
		case 2*i <= n:
			tree[i-1] = node{name: i, parent: i / 2, left: i * 2}
		default:
			tree[i-1] = node{name: i, parent: i / 2}
		}
	}

	root := tree[0]

	for _, mut := range mutations { //going through permutations
		box := tree[mut-1]

		if box.parent == 0 { // checking if the node is already a root node
			continue
		}

		boxParent := tree[tree[mut-1].parent-1]

		tree[mut-1].parent = boxParent.parent
		tree[box.parent-1].parent = box.name

		if tree[box.parent-1].left == box.name { //checking if the permutating node is on the left of its parent node
			tree[box.parent-1].left = box.left
			tree[mut-1].left = boxParent.name

			if box.left != 0 {
				tree[box.left-1].parent = boxParent.name
			}

		} else {
			tree[box.parent-1].right = box.right
			tree[mut-1].right = boxParent.name

			if box.right != 0 {
				tree[box.right-1].parent = boxParent.name
			}
		}

		if boxParent.parent != 0 { // checking if the parent node is a root node
			if tree[boxParent.parent-1].left == boxParent.name {
				tree[boxParent.parent-1].left = box.name
			} else {
				tree[boxParent.parent-1].right = box.name
			}
		} else {
			root = tree[mut-1]
		}
	}

	fo, _ := os.Create("output.txt")
	writer := bufio.NewWriter(fo)
	tree.out(root.name, writer)
	writer.Flush()
}
