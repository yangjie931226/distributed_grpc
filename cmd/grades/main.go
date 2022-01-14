package main

import (
	"context"
	"distributed/grpc/config"
	"distributed/grpc/grades"
	"distributed/grpc/log"
	"distributed/grpc/grades/pb"
	"distributed/grpc/proto"
	"distributed/grpc/registry"
	"fmt"
	"google.golang.org/grpc"
	stlog "log"
	"net"
)

func main() {
	addr := fmt.Sprintf("%s:%d", config.GobalConfig.IP, config.GobalConfig.Port)
	//httpaddr := fmt.Sprintf("http://%s:%d", config.GobalConfig.IP, config.GobalConfig.Port)
	reg := registry.Registion{
		ServiceName: registry.ServiceName(config.GobalConfig.ServerName),
		ServiceUrl:addr,
		ServiceUpdateUrl:config.GobalConfig.ServicesUpdateUrl,
		RequiresService:[]registry.ServiceName{
			registry.LOG_SERVICE,
		},
		HeartbeatUrl:config.GobalConfig.HeartbeatUrl,

	}
	ctx,err := start(context.Background(),reg,addr)

	if err != nil {
		stlog.Println(err)
		return
	}
	if logProvider, err := registry.GetProvider(registry.LOG_SERVICE); err == nil {
		fmt.Println(233232)
		log.SetLogger(config.GobalConfig.ServerName,logProvider)
		fmt.Println(2332222232)

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
	pb.RegisterGradesServiceServer(srv,&grades.GradesService{})
	proto.RegisterCommonServiceServer(srv,&log.LogService{})
	go func() {
		err := srv.Serve(listener)
		if err!= nil {
			stlog.Println(err)
		}
		err = registry.RegistryRemove(reg)
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
		err := registry.RegistryRemove(reg)
		if err!= nil {
			stlog.Println(err)
		}
	}()
	err = registry.RegistryAdd(reg)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}
