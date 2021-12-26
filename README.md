Gateway
===

API Gateway collection

## HomeBody

집에 있는시간을 저장해주는 HomeBody service



Quick start
--- 

1) docker network create
- docker network create redis-net

2) docker 설치
- docker pull redis
- docker run --rm -d --network redis-net -p 6379:6379 -v ~/tmp/redis:/data --name redis redis

3) homebody server start
- cmd/homebody/main.go 실행

4) redis-client 
- docker exec -it redis bash
- redis-cli


Data structure
---

```golang 
type Account struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Image     string   `json:"image"`
	SSID      string   `json:"ssid"`
	BSSID     string   `json:"bssid"`
	Street    string   `json:"street"`
	InitDate  string   `json:"initDate"`
	Latitude  float64  `json:"latitude"`
	Longitude float64  `json:"longitude"`
	Friends   []string `json:"friends"`
}

type Friend struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Image  string `json:"image"`
	AtHome bool   `json:"atHome"`
}

type AccountHeader struct {
	ID string `header:"id" binding:"required"`
}


type FriendHeader struct {
	ID  string `header:"id" binding:"required"`
	FID string `header:"fid" binding:"omitempty"`
}
```

REST API
---
+ account set
  + url : http://{{ endpoint }}:{{ port }}/account/set
  + method : POST
  + body: `Account`
+ account get
  + url : http://{{ endpoint }}:{{ port }}/account/get  
  + method : GET
  + header: `getAccountHeader`
+ account delete
  + url : http://{{ endpoint }}:{{ port }}/account/delete  
  + method : GET   
  + header: `AccountHeader`
+ location set
  + url : http://{{ endpoint }}:{{ port }}/location/set
  + method : POST
  + body: `Account`
+ wifi set
  + url : http://{{ endpoint }}:{{ port }}/wifi/set
  + method : POST
  + body: `Account`
+ friend set
  + url : http://{{ endpoint }}:{{ port }}/friend/set  
  + method : POST
  + header : `FriendHeader`
+ friend getAll
  + url : http://{{ endpoint }}:{{ port }}/friend/get/all
  + method : GET   
  + header : `FriendHeader`   
  + return : `[]Friend`
+ friend get
  + url : http://{{ endpoint }}:{{ port }}/friend/get
  + method : GET   
  + header : `FriendHeader`   
  + return : `Friend`
+ friend delete
  + url : http://{{ endpoint }}:{{ port }}/friend/delete
  + method : POST
  + header : `FriendHeader`

Config
---

```yaml
Web:
  port: ":8080" // web-server open port

DB:
  type: redis // temparary cache db
  address: "localhost:6379" // connection with docker

serializer:
  type: json 
```

