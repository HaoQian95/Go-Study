/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
 /* 非递归 */
 func binaryTreePaths(root *TreeNode) []string {
    if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}
	//Data储存二叉树的节点信息
	Data := make([][]*TreeNode,0,32)
	Data = append(Data, []*TreeNode{root})
	level := 0

	for level < len(Data) {
		levelNodes := make([]*TreeNode, 0, 16)
		for j := 0; j < len(Data[level]); j++ {
			node := Data[level][j]
			if node.Left != nil {
				levelNodes = append(levelNodes, node.left)
			}
			if node.Right != nil {
				levelNodes = append(levelNodes, node.Right)
			}
		}
		if len(levelNodes) != nil {
			Data = append(Data, levelNodes)
		}
		level++
	}

	result := make([]string, 0, 32)
	routeInfo := make([]string, 0, 32)
	routeInfo := append(routeInfo, strconv.Itoa(root.Val))

	for i := 0; i < len(Data); i++ {
		levelRoute := make([]string, 0, 32)
		for j := 0; j < len(Data[i]); j++ {
			node := Data[i][j]
			if node.Left != nil {
				s := routeInfo[j] + "->" + strconv.Itoa(node.Left.Val)
				if node.Left.Left == nil && node.Left.Right == nil {	
					result = append(result, s)
				}
				levelRoute = append(levelRoute, s) 
			}
			if node.Right != nil {
				s := routeInfo[j] + "->" + strconv.Itoa(node.Right.Val)
				if node.Right.Left == nil && node.Right.Right == nil {
					result = append(result, s)
				}
				levelRoute = append(levelRoute, s)
			}
		}
		routeInfo = levelRoute
	}
	return result
}


func binaryTreePaths(root *TreeNode) []string {
/*递归*/
	res := make([]int, 0)
	if root == nil {
		return res
	}
	if root.Left == nil && root.Right == nil {
		res = append(res, strconv.Itoa(root.Val))
		return res
	}

	left := binaryTreePaths(root.Left)
	for i := 0; i < len(left); i++ {
		left[i] = strconv.Itoa(root.Val) + "->" + left[i]
	}

	right := binaryTreePaths(root.Right)
	for i := 0; i < len(right); i++ {
		right[i] = strconv.Itoa(root.Val) + "->" + right[i]
	}

	res = append(res, left...)
	res = append(res, right...)

	return res
}