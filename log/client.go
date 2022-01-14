package log

import (
	"context"
	"distributed/grpc/log/pb"
	"fmt"
	"google.golang.org/grpc"
	stlog "log"
)

func SetLogger(serviceName string, serviceUrl string) {
	stlog.SetPrefix(fmt.Sprintf("[%v] - ", serviceName))
	stlog.SetFlags(0)
	stlog.SetOutput(&logWriter{url:serviceUrl})
}

type logWriter struct{
	url string
}

func (lw *logWriter) Write(p []byte) (n int, err error) {
	conn,err := grpc.Dial(lw.url, grpc.WithInsecure())
	if err != nil {
		stlog.Println(err)
		return 0, err
	}
	client := pb.NewLogServiceClient(conn)
	in := &pb.WriteLogRequest{
		Message:string(p),
	}
	_, err = client.WriteLog(context.Background(), in)
	if err != nil {
		stlog.Println(err)
		return 0, err
	}

	return len(p), nil
}
