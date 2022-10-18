package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"lxf.com/proto-demo/service"
	"sync"
	"time"
)

var (
	ProdServiceClient service.ProdServiceClient
)

type Conn struct {
	clientConn  *grpc.ClientConn
	callTimeOut time.Duration
}

func InitConnections(addr string) {
	addr = "127.0.0.1:8002"
	cc := GetConn(addr)
	ProdServiceClient = service.NewProdServiceClient(cc)
}

func GetConn(addr string) *grpc.ClientConn {
	c := &Conn{
		callTimeOut: 1 * time.Second,
	}
	if err := c.Dial(addr); err != nil {
		log.Printf("Error getting connection:%s", err.Error())
	}

	return c.clientConn
}

func (c *Conn) Dial(addr string) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.callTimeOut)
	defer cancel()

	var err error
	c.clientConn, err = grpc.DialContext(ctx, addr, grpc.WithInsecure())
	if err != nil {
		log.Printf("dial rpc error:%s", err.Error())
	}
	return nil
}

//var ch = make(chan int32, 10)

func main() {
	InitConnections("")

	var wg = sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 20; i++ {

		go func(i int) {
			//time.Sleep(time.Second)
			defer wg.Done()
			request := &service.ProductRequest{ProdId: int32(i)}
			resp, err := ProdServiceClient.GetProductStock(context.Background(), request)
			if err != nil {
				log.Printf("err:%v\n", err.Error())
			}

			fmt.Println("调用grpc方法成功11，stock=", resp.ProdStock)
		}(i)
	}
	wg.Wait()

	fmt.Println("xxx")
	time.Sleep(time.Second * 3)

	request := &service.ProductRequest{ProdId: int32(100)}
	resp, err := ProdServiceClient.GetProductStock(context.Background(), request)
	if err != nil {
		log.Printf("err:%v\n", err.Error())
	}

	fmt.Println("调用grpc方法成功12，stock=", resp.ProdStock)
}
