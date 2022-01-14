package portal

import (
	"distributed/grpc/proto"
	"distributed/grpc/registry"
	"fmt"
	"context"
)

type PortalService struct{}



func (ls *PortalService) Heatbeat(context.Context, *proto.HeatbeatRequest) (*proto.CommonReply, error) {
	return &proto.CommonReply{}, nil
}


func (ls *PortalService) UpdateProvdier(context context.Context, in *proto.PatchRequest) (*proto.CommonReply, error) {

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
