**并未使用任何原公司内部代码**
----
# 安能摧眉折腰事权贵,使我不得开心颜

发出hydra命令的客户端

# 效果图
[![效果图](https://raw.githubusercontent.com/zhangyiming748/HydraClient/master/效果图.webp)](https://raw.githubusercontent.com/zhangyiming748/HydraClient/master/效果图.webp "点击查看大图")

# 接口文档

---
title: hydra密码破解后端实现 v1.0.0
language_tabs:
- shell: Shell
- http: HTTP
- javascript: JavaScript
- ruby: Ruby
- python: Python
- php: PHP
- java: Java
- go: Go
  toc_footers: []
  includes: []
  search: true
  code_clipboard: true
  highlight_theme: darkula
  headingLevel: 2
  generator: "@tarslib/widdershins v4.0.5"

---

# hydra密码破解后端实现

> v1.0.0

# 客户端

## POST 服务端发送请求

POST /hydra/create

使用post方法实现

> Body 请求参数

```yaml
task_name: "1"
address: "2"
port: "3"
protocol: "4"
username: "5"
username_type: "6"
password: "7"
password_type: "8"
path: "9"
form: "10"
sid: "11"
username_file: file:///Users/zen/Documents/example.sh
password_file: file:///Users/zen/Documents/example.sh

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» task_name|body|string| 否 |none|
|» address|body|string| 否 |none|
|» port|body|string| 否 |none|
|» protocol|body|string| 否 |none|
|» username|body|string| 否 |none|
|» username_type|body|string| 否 |none|
|» password|body|string| 否 |none|
|» password_type|body|string| 否 |none|
|» path|body|string| 否 |none|
|» form|body|string| 否 |none|
|» sid|body|string| 否 |none|
|» username_file|body|string(binary)| 否 |none|
|» password_file|body|string(binary)| 否 |none|

> 返回示例

> 成功

```json
{
  "task_id": 33399,
  "task_name": "肯德基疯狂星期四",
  "address": "192.168.1.5",
  "port": "22",
  "protocol": "ssh",
  "username": "zen|root",
  "username_file": "",
  "user_name_type": 2,
  "password": "163453",
  "password_file": "",
  "passwd_type": 2,
  "user_id": 1,
  "path": "9",
  "form": "10",
  "sid": "11",
  "request_host": "http://127.0.0.1:2147/hydra/recv"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|
