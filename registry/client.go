package registry

import (
	"context"
	"distributed/grpc/config"
	"distributed/grpc/registry/pb"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"sync"
)

func RegistryAdd(reg Registion) error {
	//远程调用注册中心的增加服务方法
	conn, err := grpc.Dial(config.GobalConfig.RegistryServer, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return err
	}
	defer conn.Close()
	requiresService := make([]string, 0)
	for _, rs := range reg.RequiresService {
		requiresService = append(requiresService, string(rs))
	}

	client := pb.NewRegistryServiceClient(conn)
	req := &pb.RegistionRequest{
		ServiceName:      string(reg.ServiceName),
		ServiceUrl:       reg.ServiceUrl,
		RequiresService:  requiresService,
		ServiceUpdateUrl: reg.ServiceUpdateUrl,
		HeartbeatUrl:     reg.HeartbeatUrl,
	}
	_,err = client.AddRegistion(context.Background(), req)
	if err != nil {
		return fmt.Errorf("RegistryAdd err: %v\n", err)
	}

	return nil
}

func RegistryRemove(reg Registion) error {
	//远程调用注册中心的增加服务方法
	conn, err := grpc.Dial(config.GobalConfig.RegistryServer, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return err
	}
	defer conn.Close()
	requiresService := make([]string, 0)
	for _, rs := range reg.RequiresService {
		requiresService = append(requiresService, string(rs))
	}

	client := pb.NewRegistryServiceClient(conn)
	req := &pb.RegistionRequest{
		ServiceName:      string(reg.ServiceName),
		ServiceUrl:       reg.ServiceUrl,
		RequiresService:  requiresService,
		ServiceUpdateUrl: reg.ServiceUpdateUrl,
		HeartbeatUrl:     reg.HeartbeatUrl,
	}
	_,err = client.RemoveRegistion(context.Background(), req)
	if err != nil {
		return fmt.Errorf("RegistryRemove err: %v\n", err)
	}

	return nil
}



type provider struct {
	serviceProviders map[ServiceName][]string
	mutex            *sync.RWMutex
}

func (pvd *provider) update(p Patch) {
	pvd.mutex.Lock()
	defer pvd.mutex.Unlock()

	for _, add := range p.Add {
		if _, ok := pvd.serviceProviders[add.ServiceName]; ok {
			pvd.serviceProviders[add.ServiceName] = append(pvd.serviceProviders[add.ServiceName], add.ServiceUrl)
		} else {
			pvd.serviceProviders[add.ServiceName] = []string{add.ServiceUrl}
		}
	}
	for _, remove := range p.Remove {
		if services, ok := pvd.serviceProviders[remove.ServiceName]; ok {
			for index, serviceUrl := range services {
				if serviceUrl == remove.ServiceUrl {
					pvd.serviceProviders[remove.ServiceName] = append(pvd.serviceProviders[remove.ServiceName][:index],pvd.serviceProviders[remove.ServiceName][index+1:]...)
				}
			}
		}
	}
	fmt.Printf("updated provider %v \n",prov.serviceProviders)
}

func (pvd *provider) get(serviceName ServiceName) (string,error) {
	pvd.mutex.RLock()
	defer pvd.mutex.RUnlock()
	if services,ok := pvd.serviceProviders[serviceName] ; ok {
		return services[int(rand.Float32() * float32(len(services)))],nil
	}
	return "", fmt.Errorf("%v service not found", serviceName)
}

func GetProvider(name ServiceName) (string, error) {
	return prov.get(name)
}

func UpdateProvider(p Patch) {
	prov.update(p)
}

var prov = provider{
	serviceProviders: map[ServiceName][]string{},
	mutex:            &sync.RWMutex{},
}
