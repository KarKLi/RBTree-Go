package main

// color 红黑树的颜色，0为黑色，1为红色
type color int

// RBNode 红黑树的节点
type RBNode struct {
	val                 int
	left, right, parent *RBNode
	color               color
}

func (node *RBNode) reverseColor() {
	if node.color == 0 {
		node.color = 1
	} else {
		node.color = 0
	}
}

func (node *RBNode) setBlack() {
	node.color = 0
}

func (node *RBNode) setRed() {
	node.color = 1
}

func (node *RBNode) isBlack() bool {
	return node.color == 0
}

func (node *RBNode) isRed() bool {
	return node.color == 1
}

func (node *RBNode) getColor() string {
	if node.isRed() {
		return "Red"
	}
	return "Black"
}

// GetVal 获得当前红黑树节点的值
func (node *RBNode) GetVal() int {
	return node.val
}

// leftRotate 左旋函数
// 其关系为
//              |                                   |
//           Parent                                Node
//        /         \                            /       \
//      alpha      Node         =>            Parent    gamma
//               /     \                     /     \
//             beta  gamma                 alpha  beta
func (node *RBNode) leftRotate() {
	// 先获得node的parent
	parent := node.parent
	// 转移node的左儿子（beta）为parent的右儿子
	parent.right = node.left
	if parent.right != nil {
		// 设置beta的父亲为parent
		parent.right.parent = parent
	}
	// 获取parent在其父节点所在的位置
	if parent.parent == nil {
		// 说明parent是根节点
	} else if parent.parent.left == parent {
		// 说明parent是左节点
		parent.parent.left = node
	} else {
		// 说明parent是右节点
		parent.parent.right = node
	}
	// 设置node的父亲为parent的父亲
	node.parent = parent.parent
	// 设置node的左儿子为parent
	node.left = parent
	// 设置parent的父亲为node
	parent.parent = node
}

// rightRotate 右旋函数
// 其关系为
//             |                                     |
//           Parent                                 Node
//        /         \                            /       \
//      Node      alpha         =>             beta     Parent
//     /   \                                            /    \
//   beta gamma                                       gamma alpha
func (node *RBNode) rightRotate() {
	// 先获得node的parent
	parent := node.parent
	// 转移node的右儿子（gamma）为parent的左儿子
	parent.left = node.right
	if parent.left != nil {
		// 设置gamma的父亲为parent
		parent.left.parent = parent
	}
	// 获取parent在其父节点所在的位置
	if parent.parent == nil {
		// 说明parent是根节点
	} else if parent.parent.left == parent {
		// 说明parent是左节点
		parent.parent.left = node
	} else {
		// 说明parent是右节点
		parent.parent.right = node
	}
	// 设置node的父亲为parent的父亲
	node.parent = parent.parent
	// 设置node的右儿子为parent
	node.right = parent
	// 设置parent的父亲为node
	parent.parent = node
}

func (node *RBNode) grandparent() *RBNode {
	if node.parent == nil {
		return nil
	}
	return node.parent.parent
}

func (node *RBNode) uncle() *RBNode {
	if node.parent == nil {
		return nil
	}
	if node.grandparent() == nil {
		return nil
	}
	if node.grandparent().left == node.parent {
		return node.grandparent().right
	} else {
		return node.grandparent().left
	}
}

// RBInsert 向红黑树中插入一个数据
func RBInsert(root **RBNode, val int) bool {
	node := insertRBTree(root, nil, val)
	if node == nil {
		return false
	}

	// 见 https://zh.wikipedia.org/wiki/%E7%BA%A2%E9%BB%91%E6%A0%91
	var insertCase1, insertCase2, insertCase3, insertCase4, insertCase5 func(*RBNode)
	insertCase1 = func(node *RBNode) {
		if node.parent == nil {
			node.setBlack()
		} else {
			insertCase2(node)
		}
	}
	insertCase2 = func(node *RBNode) {
		if node.parent.isBlack() {
			return
		} else {
			insertCase3(node)
		}
	}
	insertCase3 = func(node *RBNode) {
		if node.uncle() != nil && node.uncle().isRed() {
			node.parent.setBlack()
			node.uncle().setBlack()
			node.grandparent().setRed()
			insertCase1(node.grandparent())
		} else {
			insertCase4(node)
		}
	}
	insertCase4 = func(node *RBNode) {
		if node == node.parent.right && node.parent == node.grandparent().left {
			node.leftRotate()
			node = node.left
		} else if node == node.parent.left && node.parent == node.grandparent().right {
			node.rightRotate()
			node = node.right
		}
		insertCase5(node)
	}
	insertCase5 = func(node *RBNode) {
		node.parent.setBlack()
		node.grandparent().setRed()
		if node == node.parent.left && node.parent == node.grandparent().left {
			node.parent.rightRotate()
		} else {
			node.parent.leftRotate()
		}
		// 如果被旋转的祖父节点是根节点，那么应该更新根节点
		if node.parent.parent == nil {
			*root = node.parent
		}
	}
	insertCase1(node)
	return true
}

func insertRBTree(root **RBNode, parent *RBNode, val int) *RBNode {
	if *root == nil {
		s := new(RBNode)
		s.val = val
		s.left = nil
		s.right = nil
		s.parent = parent
		s.setRed()
		*root = s
	} else if val < (*root).val {
		return insertRBTree(&((*root).left), *root, val)
	} else {
		return insertRBTree(&((*root).right), *root, val)
	}
	return *root
}
