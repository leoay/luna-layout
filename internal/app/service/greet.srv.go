package service

import (
	"context"
	"github.com/google/wire"

	"luna-layout/internal/app/dao"
	"luna-layout/internal/app/schema"
)

var GreetSet = wire.NewSet(wire.Struct(new(GreetSrv), "*"))

type GreetSrv struct {
	TransRepo *dao.TransRepo
	GreetRepo *dao.GreetRepo
}

func (a *GreetSrv) Query(ctx context.Context, params schema.GreetQueryParam, opts ...schema.GreetQueryOptions) (*schema.GreetQueryResult, error) {
	return a.GreetRepo.Query(ctx, params, opts...)
}

func (a *GreetSrv) Get(ctx context.Context, id uint64, opts ...schema.GreetQueryOptions) (*schema.Greet, error) {
	return nil, nil
}

func (a *GreetSrv) Create(ctx context.Context, item schema.Greet) (*schema.IDResult, error) {
	return nil, nil
}

func (a *GreetSrv) Update(ctx context.Context, id uint64, item schema.Greet) error {
	return nil
}

func (a *GreetSrv) Delete(ctx context.Context, id uint64) error {
	return nil
}
