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

const _MAX_NODE_ID int64 = 35
const _ID_COUNT_MAX_PE_MILLISECOND int64 = 36
const _INT_BASE int64 = 36

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
type Config struct {
	NodeId                  int64 // nodeId:[0,35]
	Since                   int64 // set start time,Number of milliseconds from 1970-00-00 00:00:00 to start time
	IdCountMaxPeMillisecond int64 // default 36
}

func NodeId(nodeId int64) Option {
	return func(config *Config) {
		config.NodeId = nodeId
	}
}
func NodeSince(since int64) Option {
	return func(config *Config) {
		config.Since = since
	}
}
func IdCountMaxPeMillisecond(idCountMaxPeMillisecond int64) Option {
	return func(config *Config) {
		config.IdCountMaxPeMillisecond = idCountMaxPeMillisecond
	}
}

type Option func(*Config)

// NewNode new a idNode
func NewNode(options ...Option) (*IdNode, error) {
	config := Config{}
	for _, f := range options {
		f(&config)
	}

	if config.NodeId < 0 {
		return nil, errors.New("nodeId should be greater than or equal to 0")
	}
	if config.NodeId > _MAX_NODE_ID {
		return nil, fmt.Errorf("nodeId should be less than or equal to %d", _MAX_NODE_ID)
	}
	if config.Since < 0 {
		return nil, errors.New("since time is invalid")
	}
	if config.Since > genTime() {
		return nil, errors.New("since time is invalid")
	}

	if config.IdCountMaxPeMillisecond < 0 {
		return nil, errors.New("IdCountMaxPeMillisecond should be greater than or equal to 0")
	}
	if config.IdCountMaxPeMillisecond == 0 {
		config.IdCountMaxPeMillisecond = _ID_COUNT_MAX_PE_MILLISECOND
	}
	idNode := &IdNode{
		mutex:          sync.Mutex{},
		nodeId:         config.NodeId,
		since:          config.Since,
		sequence:       0,
		lastTimestamp:  0,
		sequenceMax:    config.IdCountMaxPeMillisecond - 1,
		sequenceMaxBit: getSequenceMaxBit(config.IdCountMaxPeMillisecond),
	}
	return idNode, nil
}

func getSequenceMaxBit(idCount int64) int {
	return len(numToBHex(idCount-1, _INT_BASE))
}

// NextId generates a unique string（Contains numbers and letters） of 10 length,Each node can generate 36 ids per millisecond.
func (idNode *IdNode) NextId() (string, error) {
	idNode.mutex.Lock()
	defer idNode.mutex.Unlock()

	timestampFromSince := genTime() - idNode.since

	if timestampFromSince < idNode.lastTimestamp {
		err := errors.New("clock moved backwards")
		return "", err
	}
	if idNode.lastTimestamp == timestampFromSince {
		idNode.sequence = idNode.sequence + 1
		if idNode.sequence > idNode.sequenceMax {
			for timestampFromSince == idNode.lastTimestamp {
				timestampFromSince = genTime() - idNode.since
			}
			if timestampFromSince < idNode.lastTimestamp {
				err := errors.New("clock moved backwards")
				return "", err
			}
			idNode.sequence = 0
		}
	} else {
		idNode.sequence = 0
	}
	if timestampFromSince >= 2821109907456 { // 2821109907456 = 36^8
		return "", errors.New("time to long error")
	}
	idNode.lastTimestamp = timestampFromSince
	id := fmt.Sprintf("%08s%s%0"+strconv.Itoa(idNode.sequenceMaxBit)+"s", numToBHex(timestampFromSince, _INT_BASE), string(CHARS[idNode.nodeId]), numToBHex(idNode.sequence, _INT_BASE))
	return id, nil
}

func genTime() int64 {
	return time.Now().UnixNano() / int64(1000000)
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
