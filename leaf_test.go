package leaf

import (
	"fmt"
	"math"
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
		fmt.Println("errr", err)
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
func TestN(t *testing.T) {
	fmt.Println(math.Pow(36, 8))
	hex := numToBHex(2.821109907456e+12-1, 36)
	fmt.Println(hex)
}
