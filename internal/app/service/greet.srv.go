package service

import (
	"context"
	"github.com/google/wire"

	"github.com/leoay/luna/pkg/errors"
	"github.com/leoay/luna/pkg/util/snowflake"
	"server/internal/app/dao"
	"server/internal/app/schema"
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
	item, err := a.GreetRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}
	return item, nil
}

func (a *GreetSrv) Create(ctx context.Context, item schema.Greet) (*schema.IDResult, error) {
	item.ID = snowflake.MustID()
	err = a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		for _, rmItem := range item.GreetMenus {
			rmItem.ID = snowflake.MustID()
			rmItem.GreetID = item.ID
			err := a.GreetMenuRepo.Create(ctx, *rmItem)
			if err != nil {
				return err
			}
		}
		return a.GreetRepo.Create(ctx, item)
	})
	if err != nil {
		return nil, err
	}
	return schema.NewIDResult(item.ID), nil
}

func (a *GreetSrv) Update(ctx context.Context, id uint64, item schema.Greet) error {
	oldItem, err := a.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	} else if oldItem.Name != item.Name {
		err := a.checkName(ctx, item)
		if err != nil {
			return err
		}
	}

	item.ID = oldItem.ID
	item.Creator = oldItem.Creator
	item.CreatedAt = oldItem.CreatedAt
	err = a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		addGreetMenus, delGreetMenus := a.compareGreetMenus(ctx, oldItem.GreetMenus, item.GreetMenus)
		for _, rmitem := range addGreetMenus {
			rmitem.ID = snowflake.MustID()
			rmitem.GreetID = id
			err := a.GreetMenuRepo.Create(ctx, *rmitem)
			if err != nil {
				return err
			}
		}

		for _, rmitem := range delGreetMenus {
			err := a.GreetMenuRepo.Delete(ctx, rmitem.ID)
			if err != nil {
				return err
			}
		}

		return a.GreetRepo.Update(ctx, id, item)
	})
	if err != nil {
		return err
	}
	return nil
}

func (a *GreetSrv) Delete(ctx context.Context, id uint64) error {
	oldItem, err := a.GreetRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}
	userResult, err := a.UserRepo.Query(ctx, schema.UserQueryParam{
		PaginationParam: schema.PaginationParam{OnlyCount: true},
		GreetIDs:        []uint64{id},
	})
	if err != nil {
		return err
	} else if userResult.PageResult.Total > 0 {
		return errors.New400Response("不允许删除已经存在用户的角色")
	}
	return nil
}
