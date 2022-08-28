package main

import (
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/trip"
	"coolcar/shared/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
)

func main() {

	logger, err := server.CreateLogger()
	if err != nil {
		log.Fatalf("create zap logger error, error message: %v",err)
	}


	//注册trip服务

	/************************************将所有的grpc服务启动注册动作统一整理到一起去，因为所有的形似都差不多，只是参数不同*************************************************
	lis, err := net.Listen("tcp", ":8082")

	//由于每个服务之间都需要验证token，所以需要在各个服务启动之前进行验证
	//在grpc中，可以在启动grpc服务的时候注入拦截器，使用拦截器的方式进行类似中间件的注入一样
	//来在真正到达调用服务之前可以进行一些必要的前置操作，如验证token、参数规则校验等
	//grpc中的拦截器，是在NewServcer中传递参数,这里使用简单的拦截器ChainUnaryInterceptor
	//他的参数是type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error)
	//也就是说只要传递了一个func和上边一样就可以，由于拦截器属于公用工具，所以最好的方式就是写在公共位置shared中

	info, err := auth.InterceptorInfo("shared/auth/public.key")



	srv := grpc.NewServer(grpc.ChainUnaryInterceptor(info))

	tripSrv := &trip.Service{
		Logger:                         logger,
	}

	rentalpb.RegisterTripServiceServer(srv,tripSrv)


	logger.Sugar().Fatalf("start trip service error",zap.Error(srv.Serve(lis)))
	***********************************将所有的grpc服务启动注册动作统一整理到一起去，因为所有的形似都差不多，只是参数不同**************************************************/

	logger.Sugar().Fatalf("start trip service error",zap.Error(server.RUNGrpcServer(&server.RunGrpcServerConfig{
		NetWork:      "tcp",
		Addr:         ":8082",
		Logger:       logger,
		Name:         "trip",
		PubKey:       "shared/auth/public.key",
		RegisterFunc: func(srv *grpc.Server) {

			rentalpb.RegisterTripServiceServer(srv,&trip.Service{
				Logger:                         logger,
			})

		},
	})))

}
