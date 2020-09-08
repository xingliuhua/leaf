[English](https://github.com/xingliuhua/leaf/blob/master/README.md)
# 叶子ID

一个生成唯一ID的库，采用go语言编写。
它是雪花ID的变种，生成的唯一的ID更短，采用字符+数字组合。单个节点默认每毫秒最多生成36个ID，每个ID长度为10。

当然你可以根据自己的需求改变，调用node.SetGenerateIDRate()即可。

|每毫秒最多|ID长度|
|---|---|
|36（默认）|10|
|1296|12|
|46656|13|

相信单个节点每毫秒46656个基本能满足业务需求了。

## 背景

产品要求生成订单编号：

* 10位长度（不确定，需求随时会变）
* 字母和数字,字母不区分大小写
* 单调递增
* 唯一

很容易想到用雪花ID来实现，但是雪花是数字，而且长度比较长。

受雪花ID启发，对雪花的ID做位数的分辨做点改变:
```text
aaaaaaaa -  a  - a...
 时间戳    节点ID  序列号
```
由于是字符加数字组合，那么时间戳是8位可以用89年，而且能指定开始时间。

节点ID 0-35能满足基本需求。

序列号默认是36，也就是一毫秒最多36个订单编号，怕以后不够用，写成了动态，可以自己指定，不过这样最终的
长度也就会变长。
 
## 功能特点

* 比雪花ID更短
* 比一般的雪花ID实现更灵活，可以指定每毫秒并发量，可以指定开始时间

## 安装
go get github.com/xingliuhua/leaf
## 使用
``` go
import "github.com/xingliuhua/leaf"
```

``` go
    err, node := leaf.NewNode(0)
	if err != nil {
		return
	}
	err = node.SetGenerateIDRate(200)
	if err != nil {
		return
	}
	startTime := time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC).UnixNano() / 1000000
	err = node.SetSince(startTime)
	if err != nil {
		return
	}
	for i := 0; i < 40; i++ {
		err, id := node.NextId()
		if err != nil {
			return
		}
		fmt.Println(id)
	}
```

## 维护

[@xingliuhua](https://github.com/xingliuhua).

## 贡献

Feel free to dive in! [Open an issue] or submit PRs.
