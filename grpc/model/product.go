package model

import (
	"context"
	"grpc_study/service"
)

var ProductService = &productService{}

type productService struct{service.UnimplementedProductServiceServer}

func (p *productService) GetProductStock(ctx context.Context, req *service.ProductRequest) (*service.ProductResponse, error) {
	return &service.ProductResponse{
		ProductStock: p.GetStockById(req.ProductId),
	}, nil
}

func (p *productService) GetStockById(id int32) int32 {
	return id
}
