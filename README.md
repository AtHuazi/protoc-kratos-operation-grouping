## protoc-kratos-operation-grouping

根据 proto 文件中的注释对 kratos 生成的 API 接口进行分组。

## 安装

```bash
go install github.com/AtHuazi/protoc-kratos-operation-grouping@latest
```

## 示例

### proto 文件
```proto
syntax = "proto3";

package cart.service.v1;

import "google/api/annotations.proto";

option go_package = "api/cart/service/v1;v1";

service Cart {
  // @tags: auth,auth2
  rpc GetCart (GetCartReq) returns (GetCartReply) {}
  // comment
  // @tags:
  rpc DeleteCart (DeleteCartReq) returns (DeleteCartReply) {}

  rpc AddItem (AddItemReq) returns (AddItemReply) {}
  rpc UpdateItem (UpdateItemReq) returns (UpdateItemReply) {}
  rpc DeleteItem (DeleteItemReq) returns (DeleteItemReply) {}
}

message GetCartReq {
  int64 user_id = 1;
}

message GetCartReply {
  message Item {
    int64 item_id = 1;
    int64 quantity = 2;
  }
  repeated Item items = 1;
}

message DeleteCartReq {
  int64 user_id = 1;
}

message DeleteCartReply {

}

message AddItemReq {
  int64 user_id = 1;
  int64 item_id = 2;
  int64 quantity = 3;
}

message AddItemReply {
  message Item {
    int64 item_id = 1;
    int64 quantity = 2;
  }
  repeated Item items = 1;
}

message UpdateItemReq {
  int64 user_id = 1;
  int64 item_id = 2;
  int64 quantity = 3;
}

message UpdateItemReply {
  message Item {
    int64 item_id = 1;
    int64 quantity = 2;
  }
  repeated Item items = 1;
}

message DeleteItemReq {
  int64 user_id = 1;
  int64 item_id = 2;
}

message DeleteItemReply {
  message Item {
    int64 item_id = 1;
    int64 quantity = 2;
  }
  repeated Item items = 1;
}
```
> 给需要分组的接口添加 `@tags: xxx` 注释，多个 tag 用 `,` 分隔

### 执行命令

```bash
protoc-kratos-operation-grouping -input="./api/*/*/*/*.proto"
```

### 生成的代码

```go
package pb

var TestServerOperationGroup = map[string][]string{
	"auth": {
		"cart.service.v1.Cart.GetCart",
	},
	"auth2": {
		"cart.service.v1.Cart.GetCart",
	},
}

func GetTestServerOperationByGroup(group string) []string {
	return TestServerOperationGroup[group]
}
```
