package Greet

import (
	"context"

	"gorm.io/gorm"

	"github.com/leoay/luna/pkg/util/structure"
	"luna-layout/internal/app/dao/util"
	"luna-layout/internal/app/schema"
)

func GetGreetDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(Greet))
}

type SchemaGreet schema.Greet

func (a SchemaGreet) ToGreet() *Greet {
	item := new(Greet)
	structure.Copy(a, item)
	return item
}

type Greet struct {
	util.Model
	//添加自定义内容
}

func (a Greet) ToSchemaGreet() *schema.Greet {
	item := new(schema.Greet)
	structure.Copy(a, item)
	return item
}

type Greets []*Greet

func (a Greets) ToSchemaGreets() []*schema.Greet {
	list := make([]*schema.Greet, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaGreet()
	}
	return list
}
