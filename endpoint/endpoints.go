package endpoint

import (
	"go-app/user-service/model"
	"go-app/user-service/service"
	"go-app/user-service/utils"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

type Endpoints struct {
	GetTransactions       endpoint.Endpoint
	GetTransactionDetails endpoint.Endpoint
}

func CreateEndpoints(s service.Service, logger utils.LogService) Endpoints {
	return Endpoints{
		GetTransactions:       makeGetTransactionEndpoint(s, logger),
		GetTransactionDetails: makeGetTransactionDetailsEndpoint(s, logger),
	}
}

func makeGetTransactionDetailsEndpoint(s service.Service, logger utils.LogService) endpoint.Endpoint {
	logger.Log().Info("Initiating GetTransaction Endpoint")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//logger.LogWithContext(ctx).Info("Received GetTransactionDetails request")
		req := request.(model.RequestModel)
		//logger.LogWithContext(ctx).Debug("Received request as : " + fmt.Sprintf("%#v", req))
		transactions, err := s.GetTransactionDetails(ctx, req)
		return transactions, err
	}
}

func makeGetTransactionEndpoint(s service.Service, logger utils.LogService) endpoint.Endpoint {
	logger.Log().Info("Initiating GetTransaction Endpoint")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//logger.LogWithContext(ctx).Info("Received GetTransaction request")
		req := request.(model.RequestModel)
		//logger.LogWithContext(ctx).Debug("Received request as : " + fmt.Sprintf("%#v", req))
		transactions, err := s.GetTransactions(ctx, req)
		return transactions, err
	}
}
