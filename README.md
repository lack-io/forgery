# forgery
forgery是python库Forgerypy的go语言版本。用来伪造一些常用数据。

## 安装
go get "github.com/xingyys/forgery"

## 使用
forgery实现了forgery_py下所有功能。

### 实例
```go
package main

import (
	"github.com/xingyys/forgery"
	"fmt"
)

func main() {
	firstname, _ := forgery.FirstName()
	fmt.Printf("Get the random firstname: %s.\n", firstname)
	lastname, _ := forgery.LastName()
	fmt.Printf("Get the random lastname: %s.\n", lastname)
	fullname, _ := forgery.FullName()
	fmt.Printf("Got the random fullname: %s.\n", fullname)
}
```
### 支持的功能
[api](https://github.com/xingyys/forgery/blob/master/doc)

## 感谢

其实主要的内容都是来自于 [https://github.com/tomekwojcik/ForgeryPy](https://github.com/tomekwojcik/ForgeryPy),
个人所做的也只是将python版本的foregery_py**翻译**成golang版本的。