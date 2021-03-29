package main

import (
	"github.com/Puff-Guo/data_struct/list"
)

func main() {
	test()
	return
}

func test() {
	intVec := []int{1, 2, 3, 4, 5}

	listR := list.CreateList()
	for _, item := range intVec {
		listR.HeadInsert(item)
	}

	listR.Print()

	listR.Reverse()

	listR.Print()
}
