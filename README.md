# 環境構築

```bash
# go.modを作る
go mod init echo-get-started

# go get github.com/labstack/echo/{version}
go get github.com/labstack/echo/v4

go get github.com/labstack/echo/v4/middleware
```

# 実行

```bash
go run server.go

http server started on [::]:1323
{"time":"2022-10-09T23:04:27.753927+09:00","id":"","remote_ip":"::1","host":"localhost:1323","method":"GET","uri":"/","user_agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36","status":200,"error":"","latency":57229,"latency_human":"57.229µs","bytes_in":0,"bytes_out":11}
{"time":"2022-10-09T23:04:28.027163+09:00","id":"","remote_ip":"::1","host":"localhost:1323","method":"GET","uri":"/favicon.ico","user_agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36","status":404,"error":"code=404, message=Not Found","latency":66946,"latency_human":"66.946µs","bytes_in":0,"bytes_out":24}
```