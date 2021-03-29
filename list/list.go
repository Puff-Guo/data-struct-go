package list

type Node struct {
	Data int
	Next *Node
}

type List struct {
	Head *Node
}

func CreateList() *List {
	return &List{
		Head: &Node{
			Data: 0,
			Next: nil,
		},
	}
}

func createNode(data int) *Node {
	return &Node{
		Data: data,
		Next: nil,
	}
}

func (l *List) Reverse() {
	if nil == l.Head.Next || nil == l.Head.Next.Next {
		return
	}

	//操作第一个节点
	// var pre *Node = nil
	// var tmp *Node = nil
	// var mov *Node = nil
	// mov = l.Head.Next
	// for nil != mov {
	// 	tmp = mov.Next
	// 	if nil != tmp {
	// 		l.Head.Next = tmp
	// 	}
	// 	mov.Next = pre
	// 	pre = mov
	// 	mov = tmp
	// }

	//操作第二个节点
	var tmp *Node = nil
	var pre *Node = l.Head.Next
	var mov *Node = l.Head.Next.Next

	pre.Next = nil
	for nil != mov {
		tmp = mov.Next
		l.Head.Next = mov
		mov.Next = pre
		pre = mov
		mov = tmp
	}
}

func (l *List) HeadInsert(data int) {
	if nil != l.Head {
		pre := l.Head.Next
		newNode := createNode(data)
		newNode.Next = pre
		l.Head.Next = newNode
	}
}

func (l *List) Print() {
	if nil != l && nil != l.Head {
		tmp := l.Head.Next
		for nil != tmp {
			print(tmp.Data)
			print(" ")
			tmp = tmp.Next
		}

		println()
	}
}
