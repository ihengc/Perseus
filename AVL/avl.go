package AVL

/****************************************************************
 * @author: Ihc
 * @date: 2022/4/20 22:40
 * @description: 平衡二叉树(AVL)
 ***************************************************************/

type Node struct {
	Key    int
	Height int
	Left   *Node
	Right  *Node
}

// NewAVLTree 创建AVL树
func NewAVLTree() *Node {
	return nil
}

// Get 指定key获取节点
func Get(root *Node, key int) *Node {
	if root == nil {
		return nil
	}
	if root.Key == key {
		return root
	} else if root.Key < key {
		root = root.Right
	} else {
		root = root.Left
	}
	return Get(root, key)
}

// Put 放入元素
func Put(root **Node, key int) {
	if *root == nil {
		*root = &Node{Key: key, Height: 1}
		return
	}
	if (*root).Key < key {
		Put(&(*root).Right, key)
	}
}
