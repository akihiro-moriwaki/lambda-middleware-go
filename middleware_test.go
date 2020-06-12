package middleware

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/stretchr/testify/assert"
)

type EventHandler func(events.APIGatewayProxyRequest, string) (events.APIGatewayProxyResponse, error)

func handler(request events.APIGatewayProxyRequest, text string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       text,
		StatusCode: 200,
	}, nil
}

func testEventHandler(h EventHandler) Handler {
	return func(request events.APIGatewayProxyRequest) (interface{}, error) {
		return h(request, "world")
	}
}

func TestMiddleware(t *testing.T) {
	lambdaHandler := lambda.NewHandler(
		Middleware(
			testEventHandler(
				handler,
			),
		))
	response, err := lambdaHandler.Invoke(context.TODO(), []byte(`{}`))
	assert.NoError(t, err)
	assert.Equal(t, "{\"statusCode\":200,\"headers\":null,\"multiValueHeaders\":null,\"body\":\"world\"}", string(response))
}
