## 说明

本次代码为go中使用gin进行增删改查

## 环境

win10, go 1.20.4

## 运行代码

```
go run main.go
```

## 客户端测试

```python
import requests
import json
data ={"name": "test1", "password": "123456"}
resp = requests.post("http://127.0.0.1:8000/login", data=data)
dic = json.loads(resp.text)
token = dic["data"]["Token"]
header = {"token": token}

# resp2 = requests.get("http://127.0.0.1:8000/GetUserList", headers=header)
# print(resp2.text)


resp2 = requests.get("http://127.0.0.1:8000/UserGet/100", headers=header)
print(resp2.text)

# data1 ={"name": "test10", "password": "111111"}
# resp2 = requests.post("http://127.0.0.1:8000/AddNewUser", headers=header, data=data1)
# print(resp2.text)

# data1 ={"name": "test90", "password": "123456", "id": 2}
# resp2 = requests.post("http://127.0.0.1:8000/EditUser", headers=header, data=data)
# print(resp2.text)

# data1 ={"id": 11}
# resp2 = requests.post("http://127.0.0.1:8000/DelUser", headers=header, data=data1)
# print(resp2.text)


```

