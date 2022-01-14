package main

import (
	"context"
	"distributed/grpc/config"
	"distributed/grpc/log"
	"distributed/grpc/proto"
	"distributed/grpc/registry"
	"distributed/grpc/log/pb"
	"google.golang.org/grpc"
	stlog "log"
	"fmt"
	"net"
)

func main() {
	log.Run(config.GobalConfig.LogPath)
	addr := fmt.Sprintf("%s:%d", config.GobalConfig.IP, config.GobalConfig.Port)
	//httpaddr := fmt.Sprintf("http://%s:%d", config.GobalConfig.IP, config.GobalConfig.Port)

	reg := registry.Registion{
		ServiceName: registry.ServiceName(config.GobalConfig.ServerName),
		ServiceUrl:addr,
		ServiceUpdateUrl:config.GobalConfig.ServicesUpdateUrl,
		RequiresService:[]registry.ServiceName{
		},
		HeartbeatUrl:config.GobalConfig.HeartbeatUrl,
	}
	ctx,err := start(context.Background(), reg, addr)

	if err != nil {
		stlog.Println(err)
		return
	}

	<-ctx.Done()
}


func start(ctx context.Context, reg registry.Registion, addr string) (context.Context, error) {
	ctx, cancel := context.WithCancel(ctx)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return ctx, err
	}

	srv := grpc.NewServer()
	pb.RegisterLogServiceServer(srv,&log.LogService{})
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
