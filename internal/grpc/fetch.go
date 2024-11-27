package grpc

import (
	"context"
	main_v1 "github.com/pelicanch1k/grpcProduct-proto/proto/gen/go/proto/main"
)

func (s *serverApi) Fetch(context.Context, *main_v1.FetchRequest) (*main_v1.FetchResponse, error) {
	panic("pizda")
}
