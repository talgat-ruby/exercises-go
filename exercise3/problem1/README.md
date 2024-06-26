# Problem 1

Create **TODO** app _REST API_ with PostgreSQL. **No 3rd party libs**.

## DB

Please docker to start PostgreSQL.

```shell
docker-compose --env-file .env up --build
```

All connection info in `.env`

## Server

Please use openapi reader/editor if you familiar with it. Just copy and 
paste the content of `todo-server.yaml` for example to https://editor-next.swagger.io/.

But we will also have the Markdown version of specs below.

# Swagger TODO server 2.0
This is a sample server TODO server.

## Version: 1.0.0

### Terms of service
http://swagger.io/terms/

**Contact information:**  
talgat.s@protonmail.com

**License:** [Apache 2.0](http://www.apache.org/licenses/LICENSE-2.0.html)

[Find out more about Swagger](http://swagger.io)

**Schemes:** http

---
## todo
Everything about your TODO list

### /todo

#### GET
##### Summary

Get all todos

##### Description

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | successful operation | [ [Todo](#todo) ] |
| 400 | Invalid request |  |

#### POST
##### Summary

Add a new todo

##### Description

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| body | body | Todo object that needs to be added to the db | Yes | [TodoCreate](#todocreate) |

##### Responses

| Code | Description |
| ---- | ----------- |
| 201 | Successfuly created |
| 400 | Invalid input |
| 422 | Validation exception |

### /todo/{todoId}

#### GET
##### Summary

Get todo by ID

##### Description

Returns a single pet

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| todoId | path | ID of todo to return | Yes | integer (int) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | successful operation | [Todo](#todo) |
| 400 | Invalid ID supplied |  |
| 404 | Todo not found |  |

#### PATCH
##### Summary

Update a todo by id

##### Description

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| todoId | path | ID of todo to update | Yes | integer (int) |
| body | body | Todo object that needs to be updated | Yes | [TodoUpdate](#todoupdate) |

##### Responses

| Code | Description |
| ---- | ----------- |
| 201 | Successfuly updated |
| 400 | Invalid ID supplied |
| 404 | Todo not found |
| 422 | Validation exception |

#### DELETE
##### Summary

Delete a todo by id

##### Description

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| todoId | path | Todo id to delete | Yes | integer (int) |

##### Responses

| Code | Description |
| ---- | ----------- |
| 201 | Successfuly deleted |
| 400 | Invalid ID supplied |
| 404 | Todo not found |

---
### Models

#### Todo

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | [TodoID](#todoid) |  | No |
| description | [TodoDescription](#tododescription) |  | No |
| done | [TodoDone](#tododone) |  | No |
| createdAt | [TodoCreatedAt](#todocreatedat) |  | No |
| updatedAt | [TodoUpdatedAt](#todoupdatedat) |  | No |

#### TodoCreate

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| description | [TodoDescription](#tododescription) |  | No |

#### TodoUpdate

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| description | [TodoDescription](#tododescription) |  | No |
| done | [TodoDone](#tododone) |  | No |

#### TodoID

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| TodoID | integer |  |  |

#### TodoDescription

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| TodoDescription | string |  |  |

#### TodoDone

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| TodoDone | boolean |  |  |

#### TodoCreatedAt

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| TodoCreatedAt | string |  |  |

#### TodoUpdatedAt

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| TodoUpdatedAt | string |  |  |


