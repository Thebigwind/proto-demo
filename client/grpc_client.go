package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"lxf.com/proto-demo/service"
	"sync"
	"time"
)

func test() {
	//creds,err2 := credentials.NewServerTLSFromFile("cert/server.pem","*.lxf.com")
	//if err2 != nil{
	//	log.Fatal(err2)
	//}

	//新建链接,无认证
	conn, err := grpc.Dial("127.0.0.1:8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//conn,err := grpc.Dial("127.0.0.1:8002")
	//有认证
	//conn,err := grpc.Dial(":8002",grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatal("连接不上：", err)
	}
	defer conn.Close()

	prodClient := service.NewProdServiceClient(conn)
	//调用prodduct.pb.go中的NewProductServiceClient方法
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			request := &service.ProductRequest{ProdId: int32(i)}
			resp, err := prodClient.GetProductStock(context.Background(), request)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("调用grpc方法成功，stock=", resp.ProdStock)
		}(i)

	}
	wg.Wait()
	//

	request := &service.ProductRequest{ProdId: 233}
	resp, err := prodClient.GetProductStock(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("调用grpc方法成功，stock=", resp.ProdStock)

	fmt.Println("time start:", time.Now().UnixMilli())
	NewClient(context.Background()).GetComment(nil)
	fmt.Println("time end:", time.Now().UnixMilli())
}

type Client struct {
	Name string
	Ctx  context.Context
}

// NewBdpAiClient ...
func NewClient(ctx context.Context) *Client {
	gc := &Client{}
	gc.Name = "service_xx_ai"
	gc.Ctx = ctx
	return gc
}

type CommentInfo struct{}

func (c *Client) GetComment(param map[string]interface{}) (CommentInfo, error) {

	content := CommentInfo{}
	conn, err := grpc.Dial("127.0.0.1:8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("连接不上：", err)
	}
	defer conn.Close()

	prodClient := service.NewProdServiceClient(conn)
	request := &service.ProductRequest{ProdId: 233}
	resp, err := prodClient.GetProductStock(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("调用grpc方法成功，stock=", resp.ProdStock)

	return content, nil
}
