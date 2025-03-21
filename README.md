# framework

核心库. 提供:

- 框架库: 类型转换, MapStruct, 统一结果, 配置管理, 组件托管, 钩子机制...等.

## 类型转换

- func As{Type} ({Type}, error)
- func To{Type} {Type}

## MapStruct

- MapStruct(m map[string]any, s any, tag string) error

## 统一结果

- type StatusResult interface

```
StatusResult统一结果与错误的数据结构, 并实现与Grpc Error的转换.
因为Grpc Error Status只有Code字段, 约定StatusResult的Status/Code分别存储在高9位与低22位! 由于int32需要保留一个符号位

约定StatusResult Code取值范围:
- [0,17)         表示保留错误码! grpc内置错误码, 参考codes._maxCode
- [17,4194393)   表示业务错误码! 最大值(2^22 - 1)! 因为Grpc Code的前10位用于表示StatusResult Status

约定StatusResult Status取值范围:
- (0,511]
```

- func StatusError(status uint32, code uint32, message string, details ...string) StatusResult

## 配置管理(TODO)

## 组件托管(TODO)

## 钩子机制(TODO)

