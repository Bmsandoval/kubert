package greeting_endpoints

import (
	"context"
	"github.com/bmsandoval/kubert/api/transport_http/codecs/requests/hello_requests"
	"github.com/bmsandoval/kubert/api/transport_http/codecs/responses/hello_responses"
	"github.com/bmsandoval/kubert/library/appcontext"
	"github.com/bmsandoval/kubert/services"
	"github.com/bmsandoval/kubert/services/grpc_service"
	"github.com/go-kit/kit/endpoint"
	"log"
	"time"
)

//Validation
//Find Campaing date
func MakeGetHelloEndpoint(appCtx appcontext.Context, services services.Bundle) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(hello_requests.GetHelloRequest)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		response, err := services.GrpcSvc.GreeterClient.GetHello(ctx, &grpc_service.GetHelloRequest{})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", response)

		return hello_responses.GetHelloResponse{
			Response: response.Greetings,
		}, nil
	}
}