package leaf_test

import (
	"fmt"
	"github.com/xingliuhua/leaf/v2"
)

func ExampleIdNode_NextId() {
	node, err := leaf.NewNode(leaf.NodeId(1), leaf.IdCountMaxPeMillisecond(20))
	// since := time.Date(1992, 10, 2, 0, 0, 0, 0, time.Local).UnixNano() / 1000000
	// node, err := leaf.NewNode(leaf.NodeId(1), leaf.IdCountMaxPeMillisecond(20), leaf.NodeSince(since))
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
	// Output:
	// ksah344o10
	// ksah344o11
	// ksah344o12
	// ksah344o13
	// ksah344o14
	// ksah344o15
	// ksah344o16
	// ksah344o17
	// ksah344o18
	// ksah344o19
}
