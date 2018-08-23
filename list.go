package main

import (
    "fmt"


)
func main(){
a:=Node{data:2}
b:=new(List)
b.Append(&a)
b.Append(&a)
b.Insert(1,&a)	
c:=b.Get()
fmt.Println(c.data)

}
//节点
type Node struct {
	data int
    next *Node
}
//单向链表
type List struct {
	head *Node
	tail *Node
	size int
} 
//初始化
func(list *List) Init(){
	list.size = 0    // 此时链表是空的
	list.head = nil  // 没有头
	list.tail = nil  // 没有尾
}
func (list *List) Append(node *Node) bool {
	if node == nil {
		return false
	}
	
	if list.size==0{
		list.head=node
	}else{
		lasttail:=list.tail
		(*lasttail).next=node
	}
		// 调整尾部位置，及链表元素数量
	list.tail = node // node成为新的尾部
	list.size++      // 元素数量增加
//fmt.Println((*node).data)
	return true
	
	}
	
func(list *List)Insert(n int,node *Node)bool{
	if node == nil || n > (*list).size || (*list).size == 0 {
		return false
	}
	if n== 0 { // 直接排第一，也就领导小舅子才可以
		(*node).next = (*list).head
		(*list).head = node
		}else{
		
		t:=(*list).head
		
		for j:=1;j<n;j++{
		t=(*t).next
	}
		(*node).next=(*t).next
		(*t).next=node
	}
	(*list).size++
//fmt.Println((*node).data)
	return true
	
}

func (list *List) Get(i int) *Node {
	if i >= (*list).size {
		return nil
	}

	item := (*list).head
	for j := 0; j < i ; j++ {    // 从head数i个
		item = (*item).next
	}

	return item
}



func Del(n int)bool{
	if i >= (*list).size {
		return false
	}
	if i == 0 { // 删除头部
		node := (*list).head
		(*list).head = (*node).next
		if (*list).size == 1 { // 如果只有一个元素，那尾部也要调整
			(*list).tail = nil
		}
		}else{
			t = (*list).head
			for j=1;j<n;j++{
			t=(*t).next	
			}
			
			node:=(*t).next
			(*t).next=(*node).
		}
		
		
	
}
