# schoolfish-refresh

这是[原项目](https://github.com/HydrogenDeuterium/schoolfish)的重制版，弥补了之前犯下的一些错误，使得提交历史更加清晰明了，利于学习。本项目是毕业设计：闲鱼校园版的后端实现，基于
gin.go ，使用 redis+MySQL 作为数据库。

参考项目:

- [这个](https://gitee.com/zfkhhh/android-trading-platform)
- [这个](https://linlinjava.gitbook.io/litemall/)

## 技术选型

gin 基于 golang，是一个轻便而敏捷的 web 框架，相比 SpringBoot 更加轻量级，并且具有更高的性能。SpringBoot 在拓展性和社区支持等方面有优势，但考虑到业务范围有限，加之考虑到 JAVA
在开发速率上的缺陷，放弃选用。

MySQL 是一个轻量级的开源数据库实现，优点在于轻便简捷，应用广泛，社区支持更多，遇到问题能够容易的找到解决方法。与竞争者 PostgreSQL
相比，缺点在于性能较差，功能有限。由于项目简单，对于数据库性能和高级性能不敏感，更加关心快速完成项目目标，选用 MySQL 作为数据库选型。

redis 是一个轻量级的内存数据库，可以提高较大型数据库的读写效率。按照设计要求本项目并不需要加速缓存，但出于提高技术水平的目的还是加入了这一层内存缓存。与 MongoDB 相比，redis 较为简单，具有更高的速度和更低的内存消耗，而
Mongo 更加适合大数据量、集群的情况，本项目不考虑此类情况，选用 redis。

## api 设计

### 接口规范

REST是 Representational State Transfer(表述性状态转移)的首字母缩写。 restful api 是一套对于 api 的规范。本项目尽量实现以下规范:

- 对于某个特定的资源（用户，商品，地址，etc），分别使用以下 HTTP 方法来对应操作:
    - GET: 获取/查找
    - POST: 新建
    - DELETE: 删除
    - PUT/PATCH: 修改
    - 注意：
        - 一个 url 不一定实现了所有的方法。
        - 部分 api 不适用于 restful 的api设计，如用户管理

### 请求规范

- 请求示例:
    - GET/DELETE api_URL?params1=val1&param2=val2
    - POST/PUT/PATCH api_URL { body }

- 数组数据的分页请求参数:
    - page:  请求页码，从 1 开始。
    - limit:  每一页数量。
    - sort:  排序字段。
    - order:  升序降序，“desc” 或 “asc”
    - 注意: 并非所有返回数组的字段都实现了以上所有功能，这仅仅是规范的建议。

### 响应规范

响应格式如下:

```http request
Content-Type: application/json;charset=UTF-8

{
    code: int
    data: object|array|null
    msg: string
}
```

- 成功响应: `{code: 200,data: xxx,msg: ""}`
    - 说明响应成功（废话）。
- 失败响应: `{code: 400,data: null,msg: "xxx"}`
    - 说明前端传入的数据不合适，数量不足或不符合后端要求。
- 内部错误响应: `{code: 500,data: null,msg: ""}`
    - 内部错误，说明后端出现了考虑到但没有解决的问题。

### 登录认证

1. 前端访问登录 api，获取 token。
2. 前端将获取的 token 存放在本地。
3. 前端请求受到保护的数据时，请求携带本地 token。
4. 后端验证 token 可靠性，返回响应（成功或失败）

### 设计局限

#### 版本控制

通常 api 应当实现版本控制，以保证兼容性。如 [http://example.domain/api/v1/xxx]()。 考虑到本项目为一次性开发，暂时不考虑实现版本控制。

#### api 保护

为保证 api 不被滥用，通常 api 需要引入保护机制，例如 OAuth2。否则实际上一旦开发者知道服务器，就很容易访问 api，可能导致 api 被滥用。 考虑到本项目仅仅是一个作业项目，并不长期运行，考虑到开发周期问题，不进行 api
保护，

## api 清单

以下列出了准备实际进行实现的 api 清单，供前端开发参考。 其中，路径中带有”:“的，为路径（path）参数； 路径中“?”后的部分，为查询（query）参数； 带“#”的，要求 cookie 中有对应参数来识别用户登录；
其他所有参数都存放在请求体（request body）中。

### 认证相关

- /auth
    - [-]GET: 获取登录验证码
    - [-]POST: 登录
    - [-]DELETE: 删除登录验证码

### 用户相关

- /users
    - [-] POST: 注册
- /users/: uid
    - [-] GET: #获取用户信息
    - [-] PUT: #修改信息

### 货物相关

- /product
    - [-] GET: 获取所有
    - [-] POST: #上架
- /product/:pid
    - [-] GET: 获取信息
    - [-]PUT: #更新信息
    - [-]DELETE: #下架
- /products/users/:uid
    - [-] GET: 获取所有

### 评论相关

- /comments/:cid
    - [-]GET: 获取
    - [-]PUT: #修改
    - DELETE:#删除
- /comments/:cid/response
    - [-]GET: 获取所有回复
    - [-]POST: #新增回复
- /comments/product/:pid
    - [-]GET: 获取所有指定货物的评论
    - [-]POST: #新增评论

### 私聊相关

- /messages/:uid/users
    - GET: #获取往来用户
- /messages/:uid1/&users=uid2
    - GET: #获取与 uid2 的所有往来消息
    - POST: #向 uid2 发送消息

## 数据表设计

使用 /transplant.sql 来导入数据表结构。

- 公用部分：
    - create:datetime = CURRENT_TIMESTAMP
    - update:datetime = CURRENT_TIMESTAMP, refresh
    - delete:datetime = NULL
- 用户表: user
    - uid: uint, primary,index, inc
    - name: varchar(16)
    - email: varchar(255), unique
    - hashed: char(64)
    - avatar: varchar(255), null
    - info: varchar(255), null
    - profile: text, null
    - location: varchar(255), null
- 货物表: products
    - pid: uint, primary, index, inc
    - title: varchar(63)
    - info: varchar(2047), null
    - price: decimal(6, 2)
    - owner: uint->user(uid)
    - location: varchar(255), null
- 评论表: comments
    - cid: serial,primary
    - product: int=produce(pid)
    - commentator: uint->user.uid
    - response_to: uint->comments.cid, null
    - text: varchar(255)
- 私信表: messages
    - mid: uint, primary,index, inc
    - from: uint->user.uid
    - to: uint->user.uid
    - text: varchar(255)
- 图表：images
    - iid: uint, primary,index, inc
    - address: text, unique
    - product: uint->product.pid
