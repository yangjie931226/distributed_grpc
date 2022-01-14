package main

import (
	"context"
	"distributed/grpc/config"
	"distributed/grpc/log"
	"distributed/grpc/portal"
	"distributed/grpc/proto"
	"distributed/grpc/registry"
	"fmt"
	"google.golang.org/grpc"
	stlog "log"
	"net"
	"net/http"
)

func main() {
	err := portal.ImportTemplates()
	if err != nil {
		stlog.Fatal(err)
	}
	addr := fmt.Sprintf("%s:%d", config.GobalConfig.IP, config.GobalConfig.Port)
	reg := registry.Registion{
		ServiceName: registry.ServiceName(config.GobalConfig.ServerName),
		ServiceUrl:addr,
		ServiceUpdateUrl:config.GobalConfig.ServicesUpdateUrl,
		RequiresService:[]registry.ServiceName{
			registry.LOG_SERVICE,
			registry.GRADES_SERVICE,
		},
		HeartbeatUrl:config.GobalConfig.HeartbeatUrl,

	}

	var srv http.Server
	srv.Addr = config.GobalConfig.HttpService

	go func() {
		portal.RegisterHandlers()
		stlog.Println(srv.ListenAndServe())
	}()

	ctx,err := start(context.Background(),reg,addr)
	if err != nil {
		stlog.Println(err)
		return
	}

	if logProvider, err := registry.GetProvider(registry.LOG_SERVICE); err == nil {
		log.SetLogger(config.GobalConfig.ServerName,logProvider)

	}



	stlog.Println("测试日志服务")
	<-ctx.Done()
}


func start(ctx context.Context, reg registry.Registion, addr string) (context.Context, error) {
	ctx, cancel := context.WithCancel(ctx)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return ctx, err
	}

	srv := grpc.NewServer()
	proto.RegisterCommonServiceServer(srv,&log.LogService{})
	go func() {
		stlog.Println(srv.Serve(listener))
		err := registry.RegistryRemove(reg)
		if err!= nil {
			stlog.Println(err)
		}
		cancel()
	}()

	go func() {
		var stop string
		fmt.Printf("%v started. Press any key to stop \n", reg.ServiceName)
		fmt.Scanln(&stop)
		srv.Stop()
	}()
	err = registry.RegistryAdd(reg)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}
