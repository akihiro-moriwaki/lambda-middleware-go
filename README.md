# lambda-middleware-go

Lambda用のmiddleware

```go
func main() {
	lambda.Start(
		Middleware(
			XxxxxxxxxHandler(
				handler,
			),
		),
	)
}
```
