package hello_responses

import (
	"context"
	"encoding/json"
	"github.com/bmsandoval/kubert/library/appcontext"
	kithttp "github.com/go-kit/kit/transport/http"
	"log"
	"net/http"
)

// Errorer is implemented by all concrete response types that may contain
// errors. It allows us to change the HTTP response code without needing to
// trigger an endpoint (transport-level) error.
type Errorer interface {
	error() error
}

type GetHelloResponse struct {
	Response []string
}

func MakeGetHelloResponseEncoder(appCtx appcontext.Context) (kithttp.EncodeResponseFunc, error) {
	return func(ctx context.Context, httpResponse http.ResponseWriter, endpointResponse interface{}) error {
		response, err := json.Marshal(endpointResponse)
		if err != nil {
			log.Println(err.Error())
		}
		_, err = httpResponse.Write(response)
		return err
	}, nil
}

// EncodeError - Provide those as HTTP errors
func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("EncodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
