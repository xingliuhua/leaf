package leaf_test

import (
	"fmt"
	"github.com/xingliuhua/leaf"
	"time"
)


func ExampleIdNode_NextId() {
	startTime := time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC).UnixNano() / 1000000
	err, node := leaf.NewNode(0)
	// node.SetGenerateIDRate(200)
	node.SetSince(startTime)
	if err != nil {
		return
	}
	for i := 0; i < 15; i++ {
		fmt.Println(node.NextId())
	}
	// Output:
	// b7dm0q600
	// b7dm0q601
	// ...
	// b7dm0q60e
}
