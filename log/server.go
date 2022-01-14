package log

import (
	"context"
	"distributed/grpc/log/pb"
	"distributed/grpc/proto"
	"distributed/grpc/registry"
	"fmt"
	stlog "log"
	"os"
)

var log *stlog.Logger

type fileLog string

func (fl fileLog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0766)
	if err != nil {
		return 0, nil
	}
	defer f.Close()
	return f.Write(data)

}

func Run(name string) {
	log = stlog.New(fileLog(name), "[go]", stlog.LstdFlags)
}

func write(data string) {
	log.Printf("%v\n", data)
}

type LogService struct{}

func (ls *LogService) WriteLog(contenxt context.Context, in *pb.WriteLogRequest) (*pb.LogReply, error) {
	write(in.Message)
	fmt.Println(in.Message)
	return &pb.LogReply{}, nil
}

func (ls *LogService) Heatbeat(context.Context, *proto.HeatbeatRequest) (*proto.CommonReply, error) {
	return &proto.CommonReply{}, nil
}


func (ls *LogService) UpdateProvdier(context context.Context, in *proto.PatchRequest) (*proto.CommonReply, error) {

	p := registry.Patch{
		Add:    []registry.PatchEntry{},
		Remove: []registry.PatchEntry{},
	}

	for _,add:= range in.Add {
		p.Add = append(p.Add, registry.PatchEntry{
			ServiceName:registry.ServiceName(add.Name),
			ServiceUrl:add.Url,
		})
	}
	for _,remove:= range in.Remove {
		p.Remove = append(p.Add, registry.PatchEntry{
			ServiceName:registry.ServiceName(remove.Name),
			ServiceUrl:remove.Url,
		})
	}

	fmt.Printf("reviced update %v\n", p)
	registry.UpdateProvider(p)
	return &proto.CommonReply{}, nil
}
