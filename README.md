# 生成grpc-go文件
protoc -I protocal/ --go-grpc_out=protocal/ --go_out=protocal/ protocal/helloworld.proto