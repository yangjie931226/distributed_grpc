package main

import (
	"context"
	"distributed/grpc/config"
	"distributed/grpc/registry"
	"distributed/grpc/registry/pb"
	"fmt"
	"google.golang.org/grpc"
	stlog "log"
	"net"
	"time"
)

func main() {
	//注册handler
	registry.DoHeartbeat(30*time.Second)
	addr := fmt.Sprintf("%s:%d", config.GobalConfig.IP, config.GobalConfig.Port)

	ctx, cancel := context.WithCancel(context.Background())

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		stlog.Fatal(err)
	}

	srv := grpc.NewServer()
	pb.RegisterRegistryServiceServer(srv,&registry.RegistyServcer{})

	go func() {
		stlog.Println(srv.Serve(listener))
		cancel()
	}()

	go func() {
		fmt.Printf("%v started. Press any key to stop \n", config.GobalConfig.ServerName)
		var stop string
		fmt.Scanln(&stop)
		srv.Stop()
		cancel()
	}()


	<-ctx.Done()
}


