# GetUI-SDK

> [个推](https://www.getui.com/cn/index.html)非官方 SDK, 第四期

[![Build Status](https://travis-ci.com/gophers-coder/getui-sdk-go.svg?branch=master)](https://travis-ci.com/gophers-coder/getui-sdk-go)

## 0. Install

``` 
go get -u -v github.com/gophers-coder/getui-sdk-go/...
```

## 1. Usage

```go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gophers-coder/getui-sdk-go/getui"
)

func main() {
	account := getui.NewGetUiAccount(
		os.Getenv("APPID"),
		os.Getenv("APPSECRET"),
		os.Getenv("APPKEY"),
		os.Getenv("MASTERSECRET"))
	account.AuthSign()
	var params getui.PushSingleParams
	params.Cid = []string{"7f253f9854b1acf37264436f8deb3e37"}
	params.Title = "今天是个好日子"
	params.Text = fmt.Sprintf("时间是: %s", time.Now().Format(time.ANSIC))
	params.TransmissionContent = "消息"
	fmt.Println(account.SaveListBodyAndPushList(params))
}

```

## 2. Doc

1. 注册个推

获取：APPID、APPSECRET、APPKEY、MASTERSECRET 账户唯一标识

2. 初始化账户结构体

```` 
var account *getui.Account
account = getui.NewGetUiAccount(
    os.Getenv("APPID"),
    os.Getenv("APPSECRET"),
    os.Getenv("APPKEY"),
    os.Getenv("MASTERSECRET"))
````

3. 身份认证

``` 
account.AuthSign()
```

4. 单推

```` 
func ExamplePushSingle(cid string) {
	var params getui.PushSingleParams
	params = getui.PushSingleParams{
		Title:               "你好",
		Text:                "今天是2019.03.23",
		TransmissionContent: "好日子",
		Cid:                 []string{cid},
		AppKey:              account.AppKey,
	}
	fmt.Println(account.PushSingle(params))
}
````

5. 群推
``` 

func ExampleSaveListAndPush(cid string) {
	//"7f253f9854b1acf37264436f8deb3e37"
	var params getui.PushSingleParams
	params.Cid = []string{cid}
	params.Title = "今天是个好日子"
	params.Text = fmt.Sprintf("时间是: %s", time.Now().Format(time.ANSIC))

	content := map[string]string{}
	content["Age"] = "20"
	content["School"] = "ShangHai"
	bytes, _ := json.Marshal(content)
	params.TransmissionContent = string(bytes)
	fmt.Println(account.SaveListBodyAndPushList(params))
}
```

6. 更多用法

- [example](example/example.go)
- [restful api](http://docs.getui.com/getui/server/rest/start/)


---

@Author: 谢伟

@WeChat: wu_xiaoshen

@ZhiHu: [谢伟](https://www.zhihu.com/people/wu-xiao-shen-16/activities)

@BiliBili:[Wuxiaoshen](https://space.bilibili.com/10056291)


