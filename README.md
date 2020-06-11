# lambda-middleware-go

Lambda用のmiddleware

```
type EventHandler func(events.APIGatewayProxyRequest, string) (events.APIGatewayProxyResponse, error)

func XxxxxxxxxHandler(h EventHandler) middleware.Handler {
	return func(request events.APIGatewayProxyRequest) (interface{}, error) {
        // 共通な処理をしてEventHandlerの引数として渡す
		return h(request, "world")
	}
}

func main() {
	lambda.Start(
		middleware.Middleware(
			XxxxxxxxxHandler(
				handler,
			),
		),
	)
}
```
