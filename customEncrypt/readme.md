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

## 注：未完成错误返回