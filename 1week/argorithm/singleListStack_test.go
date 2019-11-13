package argorithm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPush(t *testing.T) {
	Head = NewStack()
	Push("GOOD")
	Push("hello")
	value := Pop()
	assert.Equal(t, value, "hello")
	value = Pop()
	fmt.Println(value,"afad")
	assert.Equal(t, value, "GOOD")
	Push("GOOD")
	Push("helloXX")
	value = Pop()
	assert.Equal(t, value, "helloXX")

	fmt.Println(Pop())
}