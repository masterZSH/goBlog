# 路由

1. 严格按照RESTFUL规范
2. 按版本进行路由分组


## 例子
获取用户信息只能GET请求

GET
host/v1/user?user_id=1

新增用户只能POST请求

POST
host/v1/user

更新用户只能PUT请求
PUT
host/v1/user/1

删除用户只能DELETE请求
DELETE
host/v1/user/1
