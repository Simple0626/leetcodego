/**
 * @author  simple
 * @date  2023/6/11 21:23
 * @version 1.0
 * 虽然在leetcode上没有通过, 但是我觉得没什么问题
 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

var head *ListNode = new(ListNode)

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			addNode(head, list1.Val)
			list1 = list1.Next
		} else {
			addNode(head, list2.Val)
			list2 = list2.Next
		}
	}
	if list1 != nil {
		addNode(head, list1.Val)
	}
	if list2 != nil {
		addNode(head, list2.Val)
	}
	return head.Next
}

func addNode(t *ListNode, v int) int {
	if head == nil {
		t = &ListNode{
			v,
			nil}
		head = t
		println("nil head")
		return 0

	}
	// 如果当前节点下一个节点为空
	if t.Next == nil {
		//t.Next指向下一个节点
		t.Next = &ListNode{
			v,
			nil}
		return -2
	}
	// 如果当前节点下一个节点不为空
	return addNode(t.Next, v)
}

func main() {

	list1Head := new(ListNode)
	addNode(list1Head, 2)
	list2Head := new(ListNode)
	addNode(list2Head, 1)
	lists := mergeTwoLists(list1Head.Next, list2Head.Next)
	for lists != nil {
		println(lists.Val)
		lists = lists.Next
	}

}
