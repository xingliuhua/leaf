package leaf

import (
	"fmt"
	"testing"
	"time"
)

func TestIdNode_NextId(t *testing.T) {
	err, node := NewNode(10)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = node.SetSince(time.Now().UnixNano()/1000000)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	err = node.SetGenerateIDRate(40)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 100; i++ {
		fmt.Println(node.NextId())
	}
}
