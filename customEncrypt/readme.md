## 三个接口

```
http://localhost:8080/encrypt/error
http://localhost:8080/encrypt/string
http://localhost:8080/encrypt/mapstring
```

## 请求体

error无请求体

string：

```json
{
  "strings": ["321", "test2"],
  "encryptionType": "md5"
}
```

mapstring：

```json
{
  "dataToEncrypt": {
    "key1": "文档",
    "key2": "test2"
  },
  "encryptionType": "md5"
}
```

## todo

1.错误未充分考虑

2.斟酌mapstring接口返回错误信息使用键还是值

