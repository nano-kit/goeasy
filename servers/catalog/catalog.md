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



### Method Catalog.List

> POST /catalog/Catalog/List <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>



Request is empty

Response parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| products | array of [object Product](#object-product) |  |


### Method Catalog.Set

> POST /catalog/Catalog/Set <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>



Request parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| product | [object Product](#object-product) |  |

Response is empty


### Method Catalog.Delete

> POST /catalog/Catalog/Delete <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>



Request parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| product_id | string |  |

Response is empty


### Method Catalog.FindByID

> POST /catalog/Catalog/FindByID <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>



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
| id | string |  |
| name | string |  |
| price_cent | int32 |  |
| snapshot | uint64 |  |
| created_at | int64 |  |
| updated_at | int64 |  |
| deleted_at | int64 |  |
| operator | string |  |

