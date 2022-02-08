package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	a := TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
	inorder(&a)
}

func inorder(root *TreeNode) []int {
	//	var stack = []int {}
	//	var nums =  [6]int {1,2,3,4,5,6}
	//	for _,v := range nums{
	//		stack = append(stack, v)    // 栈push操作
	//	}
	//	fmt.Println(stack)
	//	for len(stack) != 0 {
	//		var top = stack[len(stack)-1]    // 栈顶元素
	//		fmt.Println(top)
	//		stack = stack[:len(stack)-1]  // 栈pop操作
	//	}

	var stack []*TreeNode
	var result []int
	cur := root
	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		cur = cur.Right
	}
	return result
}
