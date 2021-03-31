package main

import (
	"github.com/Puff-Guo/data_struct/list"
	"github.com/Puff-Guo/data_struct/other/visitor"
)

func main() {
	//test()
	testVisit()
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

var count = 0

func testVisit() {
	info := visitor.Info{}
	var v visitor.Visitor = &info
	v = visitor.LogVisitor{v}
	v = visitor.NameVisitor{v}
	v = visitor.OtherThingsVisitor{v}
	loadFile := func(info *visitor.Info, err error) error {
		info.Name = "puff"
		info.Namespace = "MegaEase"
		info.OtherThings = "We are running as remote team."
		count++
		println("count:", count)
		return nil
	}

	v.Visit(loadFile)

	println("end count:", count)
}

//v.Visitor.Visitor.Visitor.Visit(loadFile)
///v.OtherThingsVisitor.Visit(NameVisitor.Visit(LogVisitor.Visit()))
