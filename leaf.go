// There are no two identical leaves in the world.
// This is a variant of Twitter snowflake. Generates a unique string，contains letters and numbers.It can be used for 89 years，
package leaf

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

const MAX_NODE_ID int64 = 35
const DEFAULT_MAX_NODE_SEQUENCE int64 = 35
const INT_BASE int64 = 36

const CHARS = "0123456789abcdefghijklmnopqrstuvwxyz"

type IdNode struct {
	mutex          sync.Mutex
	nodeId         int64
	since          int64
	sequence       int64
	lastTimestamp  int64
	sequenceMax    int64
	sequenceMaxBit int
}

// NewNode new a idNode
//
// nodeId:[0,35]
func NewNode(nodeId int64) (error, *IdNode) {
	if nodeId < 0 {
		return errors.New("nodeId should be greater than or equal to 0"), nil
	}
	if nodeId > MAX_NODE_ID {
		return fmt.Errorf("nodeId should be less than or equal to %d", MAX_NODE_ID), nil
	}

	id := &IdNode{
		mutex:          sync.Mutex{},
		nodeId:         nodeId,
		since:          0,
		sequence:       1,
		lastTimestamp:  -1,
		sequenceMax:    DEFAULT_MAX_NODE_SEQUENCE,
		sequenceMaxBit: 1,
	}

	return nil, id
}

// SetSince set start time,Number of milliseconds from 1970-00-00 00:00:00 to start time
func (i *IdNode) SetSince(sinceMilliseconds int64) error {
	if sinceMilliseconds < 0 {
		return errors.New("since time is invalid")
	}
	i.since = sinceMilliseconds
	return nil
}

// SetGenerateIDRate
func (i *IdNode) SetGenerateIDRate(idCountMaxPeMillisecond int64) error {
	if idCountMaxPeMillisecond <= 0 {
		return errors.New("sequenceMax time is invalid")
	}
	i.sequenceMax = idCountMaxPeMillisecond
	i.sequenceMaxBit = getSequenceMaxBit(idCountMaxPeMillisecond)
	return nil
}
func getSequenceMaxBit(idCount int64) int {
	return len(numToBHex(idCount, INT_BASE))
}

// NextId generates a unique string（Contains numbers and letters） of 10 length,Each node can generate 36 ids per millisecond.
func (i *IdNode) NextId() (error, string) {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	genTime := i.genTime()
	timestamp := genTime - i.since
	if timestamp < i.lastTimestamp {
		err := errors.New("clock moved backwards")
		return err, ""
	}
	if i.lastTimestamp == timestamp {
		i.sequence = i.sequence + 1
		if i.sequence > i.sequenceMax {
			return errors.New("node generates too many IDS per millisecond"), ""
		}
	} else {
		i.sequence = 1
	}
	if timestamp >= 2.821109907456e+12 {
		return errors.New("time to long error"), ""
	}
	i.lastTimestamp = timestamp
	return nil, fmt.Sprintf("%08s%s%0"+strconv.Itoa(i.sequenceMaxBit)+"s", numToBHex(timestamp, INT_BASE), string(CHARS[i.nodeId]), numToBHex(i.sequence, INT_BASE))
}

func (i *IdNode) genTime() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func numToBHex(num int64, n int64) string {
	numStr := ""
	for num != 0 {
		yu := num % n
		numStr = string(CHARS[yu]) + numStr
		num = num / n
	}
	return numStr
}
