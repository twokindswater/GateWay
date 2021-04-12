Gateway
===

API Gateway collection

## HomeBody

집에 있는시간을 저장해주는 HomeBody service



###Quick start

--- 
1) docker 설치
- docker pull redis
- docker run --rm -d --name redis-test -p 6379:6379 -v ~/tmp/redis:/data redis

2) homebody server start
- cmd/homebody/main.go 실행

###REST API 

---
[account] - set with kakao
url : http://{{ endpoint }}:{{ port }}/login/kakao
method : post
format : json
body: json {
    id string
}

> ex) http://localhost:8080/login/kakao
> method : POST
> body : { "id":"1", "image":"http://images" }


[time] - set
url : http://{{ endpoint }}:{{ port }}/time/day/set
method : post
format : json
body: json{
    id string
    date int // YYMMDD 형식 ex) 210501 : 21년 5월 1일
    time int // 오늘 집에 있었던 총 분
}

> ex) http://localhost:8080/time/day/set
> method : POST
> body : { "id":"1", "date":210501, "time":120 }


[time] - get
url : http://{{ endpoint }}:{{ port }}/time/day/get/:user/:date
method : get
format : json
return : json {
    time int // 오늘 집에 있었던 총 분
}

> ex) http://localhost:8080/time/day/get/1/210501
> method : GET

[account] - get
url : http://{{ endpoint }}:{{ port }}/account/get/:user
method : get
format : json
return : json {
    id string
    image url
    ssid string
    bssid string
}

> ex) http://localhost:8080/account/get/1
> method : GET
> Return : { "id":"1", "image":"http://images" }

###Config

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

