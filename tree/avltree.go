package tree

type elemType int32

type node struct {
	Data   elemType
	Left   *node
	Right  *node
	Parent *node
}

type tree *node

func Create() tree {
	return nil
}

//前序遍历:root->left->right
func (t *tree) FTraversal() {
	for nil != t {
		printl("%d", t.Data)
	}
}

//中序遍历:left->root->right
func (t *tree) MTraversal() {

}

//后续遍历:left->right->root
func (t *tree) BTraversal() {

}

//F:1 2 4 5 3 6 7
//M:4 2 5 1 6 3 7
//B:4 5 2 6 7 3 1
