package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	value    int
	children []*Node
}

func (n *Node) print() {
	fmt.Println(n.value)
	for _, node := range n.children {
		node.print()
	}
}

func (n *Node) hasVal(val int) bool {
	if len(n.children) == 0 {
		if val == n.value {
			return true
		}
		return false
	}
	flag := false
	for _, item := range n.children {
		if item.hasVal(val) {
			flag = true
		}
	}
	return flag
}

func main() {
	strs := readFile()
	total := 0
	for _, val := range strs {
		spl := strings.Split(val, ": ")
		nums := convAll(strings.Split(spl[1], " "))
		parent := buildTree(nums[0], nums[1:])
		vals := convToInt(spl[0])
		if parent.hasVal(vals) {
			total += vals
		}
	}
	// __AUTO_GENERATED_PRINT_VAR_START__
	fmt.Println(fmt.Sprintf("main total: %v", total)) // __AUTO_GENERATED_PRINT_VAR_END__
}

func convAll(strs []string) []int {
	var nums []int
	for _, val := range strs {
		nums = append(nums, convToInt(val))
	}
	return nums
}

func buildTree(value int, nums []int) *Node {
	if len(nums) == 0 {
		return &Node{
			value: value,
		}
	}
	return &Node{
		value: value,
		children: []*Node{
			buildTree(value*nums[0], nums[1:]),
			buildTree(value+nums[0], nums[1:]),
			buildTree(
				concatTwoNums(value, nums[0]),
				nums[1:],
			),
		},
	}
}

func concatTwoNums(a, b int) int {
	val, _ := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	return val
}
