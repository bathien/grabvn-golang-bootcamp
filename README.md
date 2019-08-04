# grabvn-golang-bootcamp
start server:
```go run apiServer.go server.go -grpc-port=8080 -db-host="YOUR_DB_HOST" -db-user="YOUR_DB_USER" -db-password="YOUR_DB_PASSWORD" -db-name="YOUR_DB_NAME"```
start client
```go run client.go -server=localhost:8080```


NOTE 8080 is configurable port so you can change to whatever you want
