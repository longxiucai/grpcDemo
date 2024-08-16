package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	// 导入grpc包
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// 导入刚才我们生成的代码所在的proto包。
	pb "grpcTest/protocal"
)

// type Lock struct {
// 	state uint32
// }

// 定义server，用来实现proto文件，里面实现的Greeter服务里面的接口
type server struct { //所有实现都必须嵌入 UnimplementedHelloServiceServer
	pb.UnimplementedGreeterServer
	// mutex         Lock
	// disableReboot bool
}

// 实现SayHello接口
// 第一个参数是上下文参数，所有接口默认都要必填
// 第二个参数是我们定义的HelloRequest消息
// 返回值是我们定义的HelloReply消息，error返回值也是必须的。
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	// 创建一个HelloReply消息，设置Message字段，然后直接返回。
	log.Println("Reques:" + in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	// 监听127.0.0.1:50051地址
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 实例化grpc服务端
	s := grpc.NewServer()

	// 注册Greeter服务
	pb.RegisterGreeterServer(s, &server{})

	// 往grpc服务端注册反射服务
	reflection.Register(s)
	log.Println("server starting")
	// 启动grpc服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
