package leaf

import (
	"fmt"
	"testing"
)

func TestIdNode_NextId(t *testing.T) {
	node, err := NewNode(NodeId(1), IdCountMaxPeMillisecond(20))
	// since := time.Date(1992, 10, 2, 0, 0, 0, 0, time.Local).UnixNano() / 1000000
	// node, err := NewNode(NodeId(1), IdCountMaxPeMillisecond(20), NodeSince(since))
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 10; i++ {
		id, err := node.NextId()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(id)
	}
}
