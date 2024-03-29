[中文版](https://github.com/xingliuhua/leaf/blob/master/README.cn.md)
# leaf

This is a variant of Twitter snowflake. Generates a unique string of 10 length，contains letters and numbers,it can be used for 89 year.
generates a unique string（Contains numbers and letters） of 10 length,Each node can generate 36 ids per millisecond.

## Background

We often need a unique ID of ten lengths,Twitter snowflake is too long.

## Feature

00000000-0-0...

It was divided into three groups：
1. timeStamp
36^8 = 2.821109907456e+12 millisecond > 89 * 365 * 24 *3600 * 1000
it can use 89 years.leaf

2. nodeId
0<= nodeId <=35

3. rate
default 36 ids per millisecond.
you can customize it.

## Install
go get github.com/xingliuhua/leaf/v2
## Usage
``` go
import "github.com/xingliuhua/leaf/v2"
```

``` go
	sinceOption := leaf.NodeSince(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC).Unix())
	node, err := leaf.NewNode(leaf.NodeId(1), leaf.IdCountMaxPeMillisecond(50), sinceOption)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 100; i++ {
		fmt.Println(node.NextId())
	}
```

## Maintainers

[@xingliuhua](https://github.com/xingliuhua).

## Contributing

Feel free to dive in! [Open an issue] or submit PRs.
