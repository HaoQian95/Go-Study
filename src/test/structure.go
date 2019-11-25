/*
二叉树实现插入排序
Author: QH
Date: 2019-11-24
*/

package test

type tree struct{
	value int
	left, right *tree
}

/*
Sort 排序
*/
func Sort(values []int){
	var root *tree
	for _, v := range values{
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}