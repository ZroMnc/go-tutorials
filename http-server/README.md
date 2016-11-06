### net/http package test

### Execute
**Build**
```bash
$ go build server.go
$ ./server
```
**Testing**
```
$ curl -v http://localhost:8081 | jq .
$ ngrep -W byline port 8081 -d lo
```

### Source:
http://soryy.com/blog/2014/not-another-go-net-http-tutorial/
