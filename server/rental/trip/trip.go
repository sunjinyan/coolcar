package trip

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/auth"
	"go.uber.org/zap"
)

type Service struct {
	Logger *zap.Logger
	rentalpb.UnimplementedTripServiceServer
}

func (s *Service) CreateTrip(ctx context.Context,req *rentalpb.CreateTripRequest) (res *rentalpb.CreateTripResponse, err error)  {

	//获取行程，首先需要获取Aid
	aid, err := auth.GetAidFromContext(ctx)
	if err != nil {
		s.Logger.Fatal("cannot get aid from context ")
	}

	s.Logger.Info("create trip ",zap.String("start",req.Start),zap.String("aid",aid))
	return &rentalpb.CreateTripResponse{}, nil
}