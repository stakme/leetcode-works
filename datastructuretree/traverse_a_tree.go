package datastructuretree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pop(list []*TreeNode) ([]*TreeNode, *TreeNode) {
	l := len(list)
	switch l {
	case 0:
		return nil, nil

	case 1:
		return nil, list[0]

	default:
		return list[:l-1], list[l-1]
	}
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := []*TreeNode{root}
	var list []int

	for {
		var node *TreeNode
		stack, node = pop(stack)
		if node == nil {
			break
		}
		list = append(list, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return list
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := []*TreeNode{}
	var list []int
	current := root
	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		stack, current = pop(stack)
		list = append(list, current.Val)
		current = current.Right
	}
	return list
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	first := []*TreeNode{root}
	var second []*TreeNode
	var current *TreeNode

	for {
		first, current = pop(first)
		if current == nil {
			break
		}
		second = append(second, current)
		if current.Left != nil {
			first = append(first, current.Left)
		}
		if current.Right != nil {
			first = append(first, current.Right)
		}
	}

	l := len(second)
	list := make([]int, l)
	for i, node := range second {
		list[l-i-1] = node.Val
	}
	return list
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var list [][]int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		var currentList []int
		var nextQueue []*TreeNode
		for _, q := range queue {
			if q.Left != nil {
				nextQueue = append(nextQueue, q.Left)
			}
			if q.Right != nil {
				nextQueue = append(nextQueue, q.Right)
			}
			currentList = append(currentList, q.Val)
		}
		list = append(list, currentList)
		queue = nextQueue
	}

	return list
}
