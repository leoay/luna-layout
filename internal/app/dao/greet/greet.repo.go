package Greet

import (
	"context"

	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/leoay/luna/pkg/errors"
	"luna-layout/internal/app/dao/util"
	"luna-layout/internal/app/schema"
)

var GreetSet = wire.NewSet(wire.Struct(new(GreetRepo), "*"))

type GreetRepo struct {
	DB *gorm.DB
}

func (a *GreetRepo) getQueryOption(opts ...schema.GreetQueryOptions) schema.GreetQueryOptions {
	var opt schema.GreetQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

func (a *GreetRepo) Query(ctx context.Context, params schema.GreetQueryParam, opts ...schema.GreetQueryOptions) (*schema.GreetQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := GetGreetDB(ctx, a.DB)

	if v := params.QueryValue; v != "" {
		v = "%" + v + "%"
		db = db.Where("Greet_name LIKE ? OR real_name LIKE ?", v, v)
	}

	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.OrderFields) > 0 {
		db = db.Order(util.ParseOrder(opt.OrderFields))
	}

	var list Greets
	pr, err := util.WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.GreetQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaGreets(),
	}
	return qr, nil
}

func (a *GreetRepo) Get(ctx context.Context, id uint64, opts ...schema.GreetQueryOptions) (*schema.Greet, error) {
	var item Greet
	ok, err := util.FindOne(ctx, GetGreetDB(ctx, a.DB).Where("id=?", id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaGreet(), nil
}

func (a *GreetRepo) Create(ctx context.Context, item schema.Greet) error {
	sitem := SchemaGreet(item)
	result := GetGreetDB(ctx, a.DB).Create(sitem.ToGreet())
	return errors.WithStack(result.Error)
}

func (a *GreetRepo) Update(ctx context.Context, id uint64, item schema.Greet) error {
	eitem := SchemaGreet(item).ToGreet()
	result := GetGreetDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

func (a *GreetRepo) Delete(ctx context.Context, id uint64) error {
	result := GetGreetDB(ctx, a.DB).Where("id=?", id).Delete(Greet{})
	return errors.WithStack(result.Error)
}

func (a *GreetRepo) UpdateStatus(ctx context.Context, id uint64, status int) error {
	result := GetGreetDB(ctx, a.DB).Where("id=?", id).Update("status", status)
	return errors.WithStack(result.Error)
}

func (a *GreetRepo) UpdatePassword(ctx context.Context, id uint64, password string) error {
	result := GetGreetDB(ctx, a.DB).Where("id=?", id).Update("password", password)
	return errors.WithStack(result.Error)
}
