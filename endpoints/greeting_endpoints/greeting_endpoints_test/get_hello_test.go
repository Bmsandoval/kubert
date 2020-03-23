package hello_test

import (
	"context"
	"github.com/bmsandoval/kubert/api/transport_http/codecs/requests/hello_requests"
	"github.com/bmsandoval/kubert/api/transport_http/codecs/responses/hello_responses"
	"github.com/bmsandoval/kubert/endpoints/greeting_endpoints"
	"github.com/bmsandoval/kubert/library/appcontext"
	"github.com/bmsandoval/kubert/mocks/services_mocks"
	"github.com/bmsandoval/kubert/services"
	"github.com/bmsandoval/kubert/services/grpc_service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetHelloServer(t *testing.T) {
	var GetHelloUnitTestData = []GetHelloTestData{
		{
			BaseTestData: BaseTestData{
				Description: "happy path",
				Request:     hello_requests.GetHelloRequest{},
				Response:    hello_responses.GetHelloResponse{
					Response: []string{"one", "two"},
				},
			},
			MockGetHello: &MockGetHello{
				InRequest: grpc_service.GetHelloRequest{},
				OutReply: grpc_service.GetHelloReply{
					Greetings:  []string{"one", "two"},
				},
				OutError:     nil,
			},
		},
	}

	_ = GetHelloUnitTestData

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, data := range GetHelloUnitTestData {
		t.Run(data.Description, func(t *testing.T) {
			serviceBundle := MockGetHelloRequiredServices(mockCtrl, data)

			f := greeting_endpoints.MakeGetHelloEndpoint(appcontext.Context{}, *serviceBundle)
			requestTestData := data.Request.(hello_requests.GetHelloRequest)
			responseTestData := data.Response.(hello_requests.GetHelloRequest)
			responseData, _ := f(context.Background(), requestTestData)
			// assert results
			//assert.Equal(t, testData.ResponseCode, res.Status.Code)

			assert.Equal(t, responseTestData, responseData)
		})
	}
}

func MockGetHelloRequiredServices(mockCtrl *gomock.Controller, data GetHelloTestData) *services.Bundle {
	greeterMock := services_mocks.NewMockGreeterClient(mockCtrl)
	greeterExpect := greeterMock.EXPECT()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if data.MockGetHello != nil {
		greeterExpect.GetHello(ctx, data.MockGetHello.InRequest).Return(
				data.MockGetHello.OutReply,
				data.MockGetHello.OutError)
	}

	serviceBundle := services.Bundle{
		GrpcSvc: grpc_service.Connection{
			Server:        nil,
			GreeterClient: greeterMock,
		},
	}

	return &serviceBundle
}
