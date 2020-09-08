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
	err = node.SetSince(time.Now().UnixNano()/1000000 + 1)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	err = node.SetGenerateIDRate(46656)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 40; i++ {
		fmt.Println(node.NextId())
	}
}
