package greeting_endpoints

import (
	"context"
	"github.com/bmsandoval/kubert/api/transport_http/codecs/requests/hello_requests"
	"github.com/bmsandoval/kubert/api/transport_http/codecs/responses/hello_responses"
	"github.com/bmsandoval/kubert/grpc/protos"
	"github.com/bmsandoval/kubert/library/appcontext"
	"github.com/bmsandoval/kubert/services"
	"github.com/go-kit/kit/endpoint"
	"log"
)

//Validation
//Find Campaing date
func MakeGetHelloEndpoint(appCtx appcontext.Context, services services.Bundle) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(hello_requests.GetHelloRequest)

		response, err := services.HelloSvc.Get(protos.HelloRequest{Name: "World"})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", response)

		return hello_responses.GetHelloResponse{
			Response: response,
		}, nil
	}
}