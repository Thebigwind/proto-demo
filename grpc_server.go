package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"lxf.com/proto-demo/service"
	"net"
)

func main() {

	//creds,err2 := credentials.NewServerTLSFromFile("cert/server.pem","cert/server.key")
	//if err2 != nil{
	//	log.Fatal(err2)
	//}

	rpcServer := grpc.NewServer()
	//rpcServer := grpc.NewServer(grpc.Creds(creds))
	service.RegisterProdServiceServer(rpcServer, service.ProdService) //RegisterProductService
	listener, err := net.Listen("tcp", "127.0.0.1:8002")
	if err != nil {
		log.Fatal(err)
	}
	err = rpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("rpc服务启动成功。。。")
}
