# schoolfish-refresh

这是[原项目](https://github.com/HydrogenDeuterium/schoolfish)的重制版，弥补了之前犯下的一些错误，使得提交历史更加清晰明了，利于学习。本项目是毕业设计：闲鱼校园版的后端实现，基于 gin.go ，使用 redis+MySQL 作为数据库。

## 技术选型

gin 基于 golang，是一个轻便而敏捷的 web 框架，相比 SpringBoot 更加轻量级，并且具有更高的性能。SpringBoot 在拓展性和社区支持等方面有优势，但考虑到业务范围有限，加之考虑到 JAVA 在开发速率上的缺陷，放弃选用。

MySQL 是一个轻量级的开源数据库实现，优点在于轻便简捷，应用广泛，社区支持更多，遇到问题能够容易的找到解决方法。与竞争者 PostgreSQL 相比，缺点在于性能较差，功能有限。由于项目简单，对于数据库性能和高级性能不敏感，更加关心快速完成项目目标，选用 MySQL 作为数据库选型。

redis 是一个轻量级的内存数据库，可以提高较大型数据库的读写效率。按照设计要求本项目并不需要加速缓存，但出于提高技术水平的目的还是加入了这一层内存缓存。与 MongoDB 相比，redis 较为简单，具有更高的速度和更低的内存消耗，而 Mongo 更加适合大数据量、集群的情况，本项目不考虑此类情况，选用 redis。

## api 设计

### 接口规范

REST是 Representational State Transfer(表述性状态转移)的首字母缩写。
restful api 是一套对于 api 的规范。本项目尽量实现以下规范:

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
