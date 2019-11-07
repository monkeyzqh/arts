package argorithm

import "fmt"

type ListNode struct {
	Value int       `json:"value"`
	Next  *ListNode `json:"next"`
}

func ReverseNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

func InitNode(arr []int) *ListNode {
	head := &ListNode{Value: 0, Next: nil}
	p := head
	for _, value := range arr {
		node := &ListNode{Value: value, Next:nil}
		p.Next = node
		p = p.Next
	}
	return head
}

func PrintNode(head *ListNode) {
	for head != nil {
		fmt.Print(head.Value)
		head = head.Next
	}
}

func main() {
	arr := []int{1,2,3,4,5}
	head := InitNode(arr)
	reverseHead := ReverseNode(head)
	PrintNode(reverseHead)
}
