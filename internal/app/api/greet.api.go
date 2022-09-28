package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"luna-layout/internal/app/ginx"
	"luna-layout/internal/app/schema"
	"luna-layout/internal/app/service"
)

var GreetSet = wire.NewSet(wire.Struct(new(GreetAPI), "*"))

type GreetAPI struct {
	GreetSrv *service.GreetSrv
}

func (a *GreetAPI) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.GreetQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.GreetSrv.Query(ctx, params, schema.GreetQueryOptions{
		OrderFields: schema.NewOrderFields(schema.NewOrderField("sequence", schema.OrderByDESC)),
	})
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResPage(c, result.Data, result.PageResult)
}

func (a *GreetAPI) QuerySelect(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.GreetQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.GreetSrv.Query(ctx, params, schema.GreetQueryOptions{
		OrderFields: schema.NewOrderFields(schema.NewOrderField("sequence", schema.OrderByDESC)),
	})
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResList(c, result.Data)
}

func (a *GreetAPI) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.GreetSrv.Get(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

func (a *GreetAPI) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Greet
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.GreetSrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

func (a *GreetAPI) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Greet
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.GreetSrv.Update(ctx, ginx.ParseParamID(c, "id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

func (a *GreetAPI) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.GreetSrv.Delete(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}
