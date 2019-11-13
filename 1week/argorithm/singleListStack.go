package argorithm

// use single list implement stack

var Head *ListNodeString
var Pre  *ListNodeString

type ListNodeString struct {
	Value string
	Next  *ListNodeString
}

func NewStack() *ListNodeString {
	head := &ListNodeString{Value: "", Next: nil}
	return head
}

func Push(value string) {
	Head.Next = &ListNodeString{Value: value, Next: nil}
	Pre = Head
	Head = Head.Next
}

func Pop() string {
	if Head == nil {
		return ""
	}
	value := Head.Value
	Head = Pre
	return value
}

func good(a int, b ...string) {

}

