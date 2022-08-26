package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net/http"
)

var Logger *zap.Logger

func InitLogger()  {
	Logger,err  := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("init logger failed,error message %v\n",err)
		return
	}
	Logger.Info("Init Logger Success")
}

func init() {
	InitLogger()
}

func main() {

	ctx,cancel :=  context.WithCancel(context.Background())
	defer cancel()
	//if err != nil {
	//	Logger.Fatal("listen port 8081 failed,error message ",zap.Error(err))
	//}

	//创建grpc gateway运行时服务，以及进行服务运行时选项配置
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard,&runtime.JSONPb{
		MarshalOptions:   protojson.MarshalOptions{
			Multiline:         false,
			Indent:            "",
			AllowPartial:      false,
			UseProtoNames:     true,
			UseEnumNumbers:    true,
			EmitUnpopulated:   true,
			Resolver:          nil,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			AllowPartial:      false,
			DiscardUnknown:    false,
			Resolver:          nil,
		},
	}))

	//注册各种服务
	err := authpb.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, "localhost:8081", []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()), //使用http，非安全连接
	}) //如果链接，使用http铭文链接，不是用https)

	if err != nil{
		Logger.Sugar().Fatalf("grpc gateway err: v%",zap.Error(err))
	}

	err = http.ListenAndServe(":8080", mux)

	if err != nil{
		Logger.Sugar().Fatalf("http service err: v%",zap.Error(err))
	}
}
