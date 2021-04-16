Gateway
===

API Gateway collection

## HomeBody

집에 있는시간을 저장해주는 HomeBody service



### Quick start

--- 
1) docker 설치
- docker pull redis
- docker run --rm -d --name redis-test -p 6379:6379 -v ~/tmp/redis:/data redis

2) homebody server start
- cmd/homebody/main.go 실행

### Data structure 

---
```golang

type AccountInfo struct {
	Id       string   `json:"id"`
	Image    string   `json:"image"`
	SSID     string   `json:"ssid"`
	BSSID    string   `json:"bssid"`
	TimeInfo TimeInfo `json:"timeInfo"`
}

type TimeInfo struct {
	Total int
	Week  int
	Day   int
}


type DayTimeInfo struct {
Id   string `json:"id"`
Date int    `json:"date"`
Time int    `json:"time"` // minute
}

```
### REST API 

---
[login] - [kakao](http://https://developers.kakao.com)  
url : http://{{ endpoint }}:{{ port }}/login/kakao   
method : post   
format : json   
body: `AccountInfo`

> ex) http://localhost:8080/login/kakao   
> method : POST   
> body : { "id":"1", "image":"http://images" }   

[login] - [facebook](https://developers.facebook.com/)   
url : http://{{ endpoint }}:{{ port }}/login/facebook   
method : post   
format : json   
body: `AccountInfo`

> ex) http://localhost:8080/login/facebook   
> method : POST   
> body : { "id":"2", "image":"http://images" }


[time] - **set**  
url : http://{{ endpoint }}:{{ port }}/time/day/set   
method : post   
format : json   
body: `DayTimeInfo`

> ex) http://localhost:8080/time/day/set   
> method : POST      
> body : { "id":"1", "date":210501, "time":120 }   


[time] - **get**   
url : http://{{ endpoint }}:{{ port }}/time/day/get/**:user**/**:date**   
method : get    
format : json    
return : json {   
    time int // 오늘 집에 있었던 총 분    
}   

> ex) http://localhost:8080/time/day/get/1/210501    
> method : GET    

[account-info] - **get**   
url : http://{{ endpoint }}:{{ port }}/account/info/get/**:user**   
method : get   
format : json   
return : `AccountInfo`   

> ex) http://localhost:8080/account/info/get/1   
> method : GET   
> Return : { "id":"1", "image":"http://images" }   

[account-AP] - **set**
url : http://{{ endpoint }}:{{ port }}/account/ap/set/**:user**   
method : post    
format : json    
body: `AccounInfo` // ssid, bssid are updated.

> ex) http://localhost:8080/account/ap/set   
> method : POST   
> Return : { "id":"1", "ssid":"asd", "bssid":"zxc" }

### Config

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

