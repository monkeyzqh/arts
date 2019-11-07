package argorithm

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestReverseNode(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	head := InitNode(arr)
	headReverse := ReverseNode(head)
	reverseValue := printNode(headReverse)

	assert.Equal(t, "543210", reverseValue)

}

func printNode(head *ListNode) string {
	reverseString := ""
	for head != nil {
		reverseString += strconv.Itoa(head.Value)
		head = head.Next
	}
	return reverseString
}
