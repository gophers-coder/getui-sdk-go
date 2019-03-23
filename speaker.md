# 演讲稿

大家好，这期我们来讲下个推（个性化推荐），是个第三方服务，通过今天的项目能让大家明白，比如经常看到的开发者文档的 SDK 等，

都是如何形成的。

比如我们之前用到的：腾讯云SDK、又拍云SDK、阿里云SDK 等。

如果你是相关公司的开发者，自己写 SDK 当然很容易，因为很容易的能访问到内部的资源。

如果不是相关公司的开发者，那么其他的开发者写 SDK 的方式是通过访问 Restful Api.

这期我们就来学习下，如何编写开源的 SDK .


## 来源

因为工作中接触到个推这块，发现相关的服务并没有 Go 版本的SDK， 所以顺手写来个 Go 版本的个推 SDK.

首先个推是什么？

我们可以看看个推服务的官网，个推实现的是个性推荐，比如推送用户感兴趣的内容，引起用户的点击行为，通过这些个性推荐的内容

我们可以研究用户的点击行为，打开程度等统计数据，进一步可以分析潜在的用户等。


## 第一步

首先我们来看下，官方给出的 Restful Api 都提供哪些功能？

- 鉴权
- 关闭鉴权
- 给用户 client id 起别名
- 给用户 client id 指定标签
- 用户黑名单
- 移除黑名单
- 单推，指定 client id 推送信息
- 群推，指定多个 client id 推送信息
- 信息又区分多重模版：点开通知、点开打开网页、点开通知弹窗、透传消息、通知栏消息布局


这些功能，一方面是为了更好的对用户管理，一方面是提供多种形态的消息推送。

看完这些文档呢，我们需要有一个意识：

- Restful 风格的 API 其实是非常常见的
- 编写符合 Restful 风格的 API 的主要考虑两方面：1： 路由设计 2：响应信息的处理 3. 设计好错误处理信息


## 第二步

我们先阅览下主要的功能和信息点：

- 鉴权：auth_token

|字段|类型|是否必须|说明|
|:---|:---|:---|:---|
|sign|string|true|sha256(appkey+timestamp+mastersecret)mastersecret为注册应用时生成|
|timestamp|string|true|时间戳，毫秒级别|
|appkey|string|true|注册应用生成的appkey|


可以看到，其实如何你想直接使用 curl 调用接口的方式来使用，发现很复杂。所以，这个时候就不如进一步封装，使用

编程语言字段的形成。

鉴权功能之后，之后的请求中都带上 auth_token 用来对用户进行鉴定操作。这个和我们常见的用户登入账户才能实现访问资源的

处理方式差不多。

处理完鉴权功能之后，后续的操作就比较简单了。无非是访问 Restful APi ，获取资源，然后对响应的信息进行处理

比如：这个起别名操作：

``` 
curl -H "Content-Type: application/json"  \
     -H "authtoken:eef0742e8bb7aa52bd1ede66a0a20c68057208656e5f558c020fb24aa5b88586" \
     https://restapi.getui.com/v1/CKWfvgBDRF8aSnGrvD7IJ4/bind_alias \
     -XPOST -d '{
                "alias_list" : [{"cid":"xxxxx", "alias":"别名"}]
                 }'
```

需要传入相应的参数，头部信息中还需要带入鉴权信息 auth_token。

我们来看下具体的代码实现


## 第三步

当然之前讲的都是些辅助的功能，是为了更好的对用户 client id 进行处理操作。

核心的功能当然是实现推送。

我们来看下：单推的功能的实现

``` 
curl -H "Content-Type: application/json" \
    -H "authtoken:f16e43e935238fffdb71c052a483267f59bbe9fac8b3cce7933b32615f75e0ef" \
     https://restapi.getui.com/v1/#{appid}/push_single \
     -XPOST -d '{
                   "message": {
                   "appkey": "appkey",
                   "is_offline": true,
                   "offline_expire_time":10000000,
                   "msgtype": "notification"
                },
                "notification": {
                    "style": {
                        "type": 0,
                        "text": "请填写通知内容",
                        "title": "请填写通知标题"
                    },
                    "transmission_type": true,
                    "transmission_content": "透传内容"
                },
                "cid": "************************",
                "requestid": "12111111111111111111111"
            }'
```

群推功能：

群推包括两个步骤：1. 保存消息体 2. 完成推送

我们来看下具体的代码实现


我们看下效果。

## 结尾

好，这节我们主要讲来下如何封装 SDK. 实现功能。