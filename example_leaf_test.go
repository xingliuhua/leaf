package leaf_test

import (
	"fmt"
	"github.com/xingliuhua/leaf"
)

func ExampleIdNode_NextId() {
	err, node := leaf.NewNode(14)
	if err != nil {
		return
	}
	// err = node.SetGenerateIDRate(200)
	// if err != nil {
	// 	return
	// }
	// startTime := time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC).UnixNano() / 1000000
	// err = node.SetSince(startTime)
	// if err != nil {
	// 	return
	// }
	for i := 0; i < 40; i++ {
		err, id := node.NextId()
		if err != nil {
			return
		}
		fmt.Println(id)
	}
	// Output:
	// 0b8l5zmf001
	// 0b8l5zmf002
	// ...
	// 0b8l5zmf014
}
