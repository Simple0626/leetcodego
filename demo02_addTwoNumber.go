package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//需要返回一个链表，这里就要创建一个链表
	var ret = new(ListNode)
	//结果链表的头节点
	current := ret
	//进位
	carry := 0

	for l1 != nil || l2 != nil {
		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		//两个数加起来加上进位
		sum := x + y + carry

		//sum % 10取sum的个位数放到结果列表中
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
		//有进位就获取进位
		carry = sum / 10
	}
	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}
	return ret.Next
}

func main() {

}
