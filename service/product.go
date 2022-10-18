package service

import "context"

var ProdService = &ProductService{}

type ProductService struct {
}

func (p ProductService) GetProductStock(context context.Context, reqeust *ProductRequest) (*ProductResponse, error) {
	//业务逻辑
	stock := p.GetStockById(reqeust.ProdId)

	return &ProductResponse{ProdStock: stock}, nil
}

func (p ProductService) GetStockById(id int32) int32 {
	return id
}

func (p ProductService) mustEmbedUnimplementedProdServiceServer() {}
