package trip

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	Logger *zap.Logger
	rentalpb.UnimplementedTripServiceServer
}


func (s *Service) CreateTrip(context.Context, *rentalpb.CreateTripRequest) (*rentalpb.TripEntity, error){

	return nil, status.Error(codes.Unimplemented,"")
}
func (s *Service) GetTrip(context.Context, *rentalpb.GetTripRequest) (*rentalpb.Trip, error){

	return nil, status.Error(codes.Unimplemented,"")
}
func (s *Service) GetTrips(context.Context, *rentalpb.GetTripsRequest) (*rentalpb.TripsRecord, error){

	return nil, status.Error(codes.Unimplemented,"")
}
func (s *Service) UpdateTrip(context.Context, *rentalpb.UpdateTripsRequest) (*rentalpb.Trip, error){

	return nil, status.Error(codes.Unimplemented,"")
}