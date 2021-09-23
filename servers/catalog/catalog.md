# API Protocol

Table of Contents

* [Service Catalog](#service-catalog)
    * [Method Catalog.List](#method-cataloglist)
    * [Method Catalog.Set](#method-catalogset)
    * [Method Catalog.Delete](#method-catalogdelete)
    * [Method Catalog.FindByID](#method-catalogfindbyid)
* [Enums](#enums)
* [Objects](#objects)
    * [Object Product](#object-product)




## Service Catalog

产品目录服务

### Method Catalog.List

> POST /catalog/Catalog/List <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>

列出所有的产品

Request is empty

Response parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| products | array of [object Product](#object-product) |  |


### Method Catalog.Set

> POST /catalog/Catalog/Set <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>

增加或者更新产品信息

Request parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| product | [object Product](#object-product) |  |

Response is empty


### Method Catalog.Delete

> POST /catalog/Catalog/Delete <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>

删除产品

Request parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| product_id | string |  |

Response is empty


### Method Catalog.FindByID

> POST /catalog/Catalog/FindByID <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>

根据产品 ID 查询

Request parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| product_ids | array of string |  |

Response parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| products | array of [object Product](#object-product) |  |





## Enums

## Objects

### object Product



Attributes

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| id | string | 产品编号 |
| name | string | 产品名称 |
| price_cent | int32 | 产品价格（单位：分） |
| snapshot | uint64 | 产品快照编号：记录历史时间线上的一个产品，被订单所引用。 |
| created_at | int64 | 创建时间 |
| updated_at | int64 | 更改时间 |
| deleted_at | int64 | 删除时间 |
| operator | string | 操作人 |

