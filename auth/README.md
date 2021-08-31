# auth

认证服务，存储和管理账号、规则(RBAC)。

## 查看权限规则

    micro auth list rules

## 查看账号

    micro auth list accounts

## 增加权限规则

缺省的规则，允许匿名请求。为了关闭缺省的规则，我们增加一条规则。

    micro auth create rule --scope '' --priority 1 --resource '*:*:*' --access denied deny-public

增加一条规则，允许具名用户请求，访问范围是 normal。访问范围（Scope）类似用户组，一个用户可以有多个访问范围（Scope）。用参数 --priority 指定更大的优先级。

    micro auth create rule --scope normal --priority 1000 --resource '*:*:*' normal-any

## 增加用户


创建一个新用户，ID 是 user001，设定其访问范围是 normal。

    micro auth create account --secret 123456 --scopes normal user001

创建成功后，获取用户 user001 的访问令牌。

    micro token --secret 123456 user001

## 参考教程

* [远程调用（RPC）的认证授权](https://nano-kit.github.io/go-micro-in-action/rpc-auth.html)
