package middleware

import (
	"github.com/aws/aws-lambda-go/events"
)

type Handler func(events.APIGatewayProxyRequest) (interface{}, error)

func Middleware(next Handler) Handler {
	return Handler(func(request events.APIGatewayProxyRequest) (interface{}, error) {
		return next(request)
	})
}
