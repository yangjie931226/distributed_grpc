package registry

import (
	"context"
	"distributed/grpc/proto"
	"distributed/grpc/registry/pb"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"sync"
	"time"
)

//管理总服务和集
type Registry struct {
	registions []Registion
	mutex      sync.RWMutex
}

func (r *Registry) add(reg Registion) error {
	r.mutex.Lock()
	r.registions = append(r.registions, reg)
	r.mutex.Unlock()
	//服务本身的依赖发给自己
	err := r.sendRequiresRegistion(reg)
	//通知其他依赖自己的服务更新依赖集和
	r.notifyAdd(reg)
	fmt.Println("Registry add", r.registions)
	return err
}

func (r *Registry) sendRequiresRegistion(reg Registion) error {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	p := Patch{
		Add:    []PatchEntry{},
		Remove: []PatchEntry{},
	}
	for _, regServiceName := range reg.RequiresService {
		for _, regionstion := range r.registions {
			if regionstion.ServiceName == regServiceName {
				pe := PatchEntry{
					ServiceName: regionstion.ServiceName,
					ServiceUrl:  regionstion.ServiceUrl,
				}
				p.Add = append(p.Add, pe)
			}
		}
	}

	//调用远程服务通知服发现
	err := r.sendPatch(p, reg.ServiceUrl)
	return err
}

func (r *Registry) notifyAdd(reg Registion) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, regionstion := range r.registions {
		//并发通知所有在线服务
		go func(regionstion Registion) {
			for _, rs := range regionstion.RequiresService {
				if rs == reg.ServiceName {
					p := Patch{
						Add:    []PatchEntry{},
						Remove: []PatchEntry{},
					}
					pe := PatchEntry{
						ServiceName: reg.ServiceName,
						ServiceUrl:  reg.ServiceUrl,
					}
					p.Add = append(p.Add, pe)
					err := r.sendPatch(p, regionstion.ServiceUrl)
					if err != nil {
						log.Println(err)
						return
					}
				}
			}

		}(regionstion)

	}

}

func (r *Registry) notifyRemove(reg Registion) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, regionstion := range r.registions {
		//并发通知所有在线服务
		go func(regionstion Registion) {
			for _, rs := range regionstion.RequiresService {
				if rs == reg.ServiceName {
					p := Patch{
						Add:    []PatchEntry{},
						Remove: []PatchEntry{},
					}
					pe := PatchEntry{
						ServiceName: reg.ServiceName,
						ServiceUrl:  reg.ServiceUrl,
					}
					p.Remove = append(p.Remove, pe)
					err := r.sendPatch(p, regionstion.ServiceUrl)
					if err != nil {
						log.Println(err)
						return
					}
				}
			}

		}(regionstion)

	}

}

func (r *Registry) sendPatch(p Patch, url string) error {

	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := proto.NewCommonServiceClient(conn)
	req := &proto.PatchRequest{
		Add:[]*proto.PatchRequest_Patch{},
		Remove:[]*proto.PatchRequest_Patch{},
	}
	for _,patchEntry := range p.Add {
		req.Add = append(req.Add, &proto.PatchRequest_Patch{
			Name:string(patchEntry.ServiceName),
			Url:patchEntry.ServiceUrl,
		})
	}
	for _,patchEntry := range p.Remove {
		req.Remove = append(req.Remove, &proto.PatchRequest_Patch{
			Name:string(patchEntry.ServiceName),
			Url:patchEntry.ServiceUrl,
		})
	}
	_, err = client.UpdateProvdier(context.Background(), req)
	//buf := new(bytes.Buffer)
	//enc := json.NewEncoder(buf)
	//err := enc.Encode(p)
	//if err != nil {
	//	return fmt.Errorf("Encode Patch error: %v", err)
	//}
	//
	//_, err = http.Post(url, "application/json", buf)
	//if err != nil {
	//	fmt.Println(err)
	//	return err
	//}
	return err
}
func (r *Registry) remove(reg Registion) error {

	for key, value := range r.registions {
		if value.ServiceUrl == reg.ServiceUrl {
			r.mutex.Lock()
			r.registions = append(r.registions[:key], r.registions[key+1:]...)
			r.mutex.Unlock()
			r.notifyRemove(reg)

			fmt.Println("Registry remove", r.registions)
			return nil
		}
	}
	//通知其他依赖自己的服务更新依赖集和
	return fmt.Errorf("service name: %v, service url: %v not found", reg.ServiceName, reg.ServiceUrl)
}
func (r *Registry) heartbeat(duration time.Duration) {
	for {

		var wg sync.WaitGroup

		for _, v := range r.registions {
			wg.Add(1)
			go func() {
				defer wg.Done()
			}()
			passFlag := true
			for i := 0; i < 3; i++ {
				conn, err := grpc.Dial(v.HeartbeatUrl, grpc.WithInsecure())
				if err != nil {
					log.Println(err)
					conn.Close()
				} else {
					client := proto.NewCommonServiceClient(conn)
					_, err = client.Heatbeat(context.Background(), &proto.HeatbeatRequest{})
					if err == nil {
						fmt.Printf("%v %v heartbeat pass \n", v.ServiceName, v.ServiceUrl)
						if !passFlag {
							r.add(v)
						}
						conn.Close()
						break
					}
				}

				if passFlag {
					passFlag = false
					r.remove(v)
				}
				time.Sleep(time.Second)

			}

		}
		wg.Wait()
		time.Sleep(duration)

	}

}

var reg = Registry{
	registions: make([]Registion, 0),
	mutex:      sync.RWMutex{},
}

type RegistyServcer struct {
}

func (rs *RegistyServcer) AddRegistion(context context.Context, in *pb.RegistionRequest) (out *pb.RegistionReply, err error) {
	out = &pb.RegistionReply{}
	out.Code = 200;
	out.Msg = ""
	requiresService := make([]ServiceName, 0)
	for _, rs := range in.RequiresService {
		requiresService = append(requiresService, ServiceName(rs))
	}
	registion := Registion{
		ServiceName:      ServiceName(in.ServiceName),
		ServiceUrl:       in.ServiceUrl,
		RequiresService:  requiresService,
		ServiceUpdateUrl: in.ServiceUpdateUrl,
		HeartbeatUrl:     in.HeartbeatUrl,
	}
	err = reg.add(registion)
	if err != nil {
		out.Code = -1;
		out.Msg = err.Error()
	}
	return
}

func (rs *RegistyServcer) RemoveRegistion(context context.Context, in *pb.RegistionRequest) (out *pb.RegistionReply, err error) {
	out = &pb.RegistionReply{}
	out.Code = 200;
	out.Msg = ""
	requiresService := make([]ServiceName, 0)
	for _, rs := range in.RequiresService {
		requiresService = append(requiresService, ServiceName(rs))
	}
	registion := Registion{
		ServiceName:      ServiceName(in.ServiceName),
		ServiceUrl:       in.ServiceUrl,
		RequiresService:  requiresService,
		ServiceUpdateUrl: in.ServiceUpdateUrl,
		HeartbeatUrl:     in.HeartbeatUrl,
	}
	err = reg.remove(registion)
	if err != nil {
		out.Code = -1;
		out.Msg = err.Error()
	}
	return
}

var once = sync.Once{}

func DoHeartbeat(duration time.Duration) {
	once.Do(func() {
		go reg.heartbeat(duration)
	})
}
