package grades

import (
	"context"
	"distributed/grpc/grades/pb"
	"distributed/grpc/proto"
	"distributed/grpc/registry"
	"fmt"
)



type GradesService struct {
}

func (gs *GradesService) GetAllStudents(ctx context.Context, in *pb.GetAllRequest) (*pb.StudentsReply, error) {
	data := []*pb.Student{}
	for _, stu := range students {
		dataGrades := []*pb.Grade{}
		for _, stuGrades := range stu.Grades {
			dataGrade := &pb.Grade{
				Title: stuGrades.Title,
				Type:  string(stuGrades.Type),
				Score: stuGrades.Score,
			}
			dataGrades = append(dataGrades, dataGrade)
		}
		dataStu := &pb.Student{
			Id:        int32(stu.ID),
			FirstName: stu.FirstName,
			LastName:  stu.LastName,
			Grades:    dataGrades,
		}
		data = append(data, dataStu)
	}
	out := &pb.StudentsReply{}
	out.Code = 200
	out.Msg = ""
	out.Data = data
	return out, nil
}

func (gs *GradesService) GetOneStudent(ctx context.Context, in *pb.IdRequest) (*pb.StudentReply, error) {
	id := int(in.Id)
	stu, err := students.GetByID(id)
	if err != nil {
		return &pb.StudentReply{}, err
	}
	dataGrades := []*pb.Grade{}
	for _, stuGrades := range stu.Grades {
		dataGrade := &pb.Grade{
			Title: stuGrades.Title,
			Type:  string(stuGrades.Type),
			Score: stuGrades.Score,
		}
		dataGrades = append(dataGrades, dataGrade)
	}
	data := &pb.Student{
		Id:        int32(stu.ID),
		FirstName: stu.FirstName,
		LastName:  stu.LastName,
		Grades:    dataGrades,
	}
	resp := &pb.StudentReply{
		Code: 200,
		Msg:  "",
		Data: data,
	}
	return resp, nil

}

func (gs *GradesService) AddGrade(ctx context.Context, in *pb.GradeRequest) (*pb.GradesReply, error) {
	id := int(in.Id)
	stu, err := students.GetByID(id)
	if err != nil {
		return &pb.GradesReply{}, err
	}

	g := Grade{
		Title: in.Grade.Title,
		Type:  GradeType(in.Grade.Type),
		Score: in.Grade.Score,
	}
	gResp := &pb.Grade{
		Title: in.Grade.Title,
		Type:  in.Grade.Type,
		Score: in.Grade.Score,
	}
	stu.Grades = append(stu.Grades, g)
	resp := &pb.GradesReply{
		Code: 200,
		Msg:  "",
		Data: gResp,
	}
	return resp, nil
}

func (gs *GradesService) Heatbeat(context.Context, *proto.HeatbeatRequest) (*proto.CommonReply, error) {
	return &proto.CommonReply{}, nil
}

func (gs *GradesService) UpdateProvdier(context context.Context, in *proto.PatchRequest) (*proto.CommonReply, error) {

	p := registry.Patch{
		Add:    []registry.PatchEntry{},
		Remove: []registry.PatchEntry{},
	}

	for _, add := range in.Add {
		p.Add = append(p.Add, registry.PatchEntry{
			ServiceName: registry.ServiceName(add.Name),
			ServiceUrl:  add.Url,
		})
	}
	for _, remove := range in.Remove {
		p.Remove = append(p.Add, registry.PatchEntry{
			ServiceName: registry.ServiceName(remove.Name),
			ServiceUrl:  remove.Url,
		})
	}

	fmt.Printf("reviced update %v \n", p)
	registry.UpdateProvider(p)
	return &proto.CommonReply{}, nil
}
