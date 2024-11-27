package grpc

import (
	"context"
	main_v1 "github.com/pelicanch1k/grpcProduct-proto/proto/gen/go/proto/main"
)

func (s *serverApi) List(context.Context, *main_v1.ListRequest) (*main_v1.ListResponse, error) {
	product := make([]*main_v1.Product, 0)
	product = append(product, &main_v1.Product{
		Id:    "1",
		Name:  "gg",
		Price: 12,
	})

	return &main_v1.ListResponse{
		Products:   product,
		TotalCount: 1,
	}, nil
}
