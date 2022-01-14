package service

import (
	"context"
	"distributed/grpc/registry"
	"distributed/grpc/registry/pb"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Start(ctx context.Context, reg registry.Registion, registyHandlerFunc func(),addr string) (context.Context, error) {
	registyHandlerFunc()
	ctx, cancel := context.WithCancel(ctx)


	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return ctx, err
	}

	srv := grpc.NewServer()
	pb.RegisterRegistryServiceServer(srv,&registry.RegistyServcer{})
	go func() {
		log.Println(srv.Serve(listener))
		err := registry.RegistryRemove(reg)
		if err!= nil {
			log.Println(err)
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
			log.Println(err)
		}
	}()
	err = registry.RegistryAdd(reg)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}
