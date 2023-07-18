package server

import (
	"context"
	"encoding/json"
	"go-app/user-service/endpoint"
	"go-app/user-service/model"
	"go-app/user-service/utils"
	"net/http"

	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var logger utils.LogService = *utils.InitLogger()

func NewGorillaMuxServer(endpoints endpoint.Endpoints) http.Handler {

	router := mux.NewRouter()

	router.Use(commonMiddleware)

	subRouter := router.PathPrefix("/transactions").Subrouter()

	subRouter.Methods("GET").Path("/account/{acc_id}").Handler(getTransactionsTransport(endpoints))

	subRouter.Methods("GET").Path("/{tx_id}").Handler(getTransactionDetailsTransport(endpoints))

	return router

}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Log().Debug("Generating correlationId")
		correlationId := uuid.New().String()
		w.Header().Add("CorrelationID", correlationId)
		ctx := context.WithValue(r.Context(), "correlationId", correlationId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func getTransactionsTransport(endpoints endpoint.Endpoints) *httpTransport.Server {
	return httpTransport.NewServer(
		endpoints.GetTransactions,
		decodeRequest,
		encodeResponse,
	)
}

func getTransactionDetailsTransport(endpoints endpoint.Endpoints) *httpTransport.Server {
	return httpTransport.NewServer(
		endpoints.GetTransactionDetails,
		decodeRequest,
		encodeResponse,
	)
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	return model.RequestModel{
		AccId:         vars["acc_id"],
		TransactionId: vars["tx_id"],
		FromDate:      r.URL.Query().Get("fromDateTime"),
		ToDate:        r.URL.Query().Get("toDateTime"),
		PageSize:      r.URL.Query().Get("page"),
		Page:          r.URL.Query().Get("pageSize"),
		Status:        r.URL.Query().Get("status"),
	}, nil
}
