package leaf

import (
	"fmt"
	"math"
	"testing"
)

func TestIdNode_NextId(t *testing.T) {
	err, node :=NewNode(0)
	node.SetGenerateIDRate(46656)
	if err != nil {
		return
	}
	for i := 0; i < 15; i++ {
		fmt.Println(node.NextId())
	}
}
func TestN(t *testing.T)  {
	fmt.Println(math.Pow(36,8))
	hex := numToBHex(2.821109907456e+12-1, 36)
	fmt.Println(hex)
}

