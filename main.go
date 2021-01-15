package main

import "fmt"

func main() {
	root := &RBNode{val: 13}
	root.setBlack()
	InputArray := []int{8, 11, 1, 6, 17, 25, 15, 22, 27}
	for _, val := range InputArray {
		RBInsert(&root, val)
	}
	// 前序遍历节点并输出它的颜色
	fmt.Println("该红黑树前序遍历为：")
	PreOrder(root)
	// 中序遍历节点并输出它的颜色
	fmt.Println("该红黑树中序遍历为：")
	MidOrder(root)
	fmt.Println("根据前序遍历和中序遍历，可以还原一棵唯一的二叉树。还原过程留作习题。")
	fmt.Println("并且请证明，对一棵红黑树的中序遍历一定是有序的。")
	return
}

// PreOrder 对二叉树进行前序遍历
func PreOrder(root *RBNode) {
	// 前序遍历是根左右的顺序
	if root != nil {
		fmt.Println("节点的值：" + fmt.Sprint(root.val) + " " + "颜色： " + root.getColor())
		PreOrder(root.left)
		PreOrder(root.right)
	}

}

// MidOrder 对二叉树进行先序遍历
func MidOrder(root *RBNode) {
	// 中序遍历是左根右的顺序
	if root != nil {
		MidOrder(root.left)
		fmt.Println("节点的值：" + fmt.Sprint(root.val) + " " + "颜色： " + root.getColor())
		MidOrder(root.right)
	}
}
