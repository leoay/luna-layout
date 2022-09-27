package schema

import (
	"fmt"
	"time"
)

// Greet 角色对象
type Greet struct {
	ID         uint64     `json:"id,string"`                             // 唯一标识
	Name       string     `json:"name" binding:"required"`               // 角色名称
	Sequence   int        `json:"sequence"`                              // 排序值
	Memo       string     `json:"memo"`                                  // 备注
	Status     int        `json:"status" binding:"required,max=2,min=1"` // 状态(1:启用 2:禁用)
	Creator    uint64     `json:"creator"`                               // 创建者
	CreatedAt  time.Time  `json:"created_at"`                            // 创建时间
	UpdatedAt  time.Time  `json:"updated_at"`                            // 更新时间
	GreetMenus GreetMenus `json:"Greet_menus" binding:"required,gt=0"`   // 角色菜单列表
}

// GreetQueryParam 查询条件
type GreetQueryParam struct {
	PaginationParam
	IDs        []uint64 `form:"-"`          // 唯一标识列表
	Name       string   `form:"-"`          // 角色名称
	QueryValue string   `form:"queryValue"` // 模糊查询
	Status     int      `form:"status"`     // 状态(1:启用 2:禁用)
}

// GreetQueryOptions 查询可选参数项
type GreetQueryOptions struct {
	OrderFields  []*OrderField // 排序字段
	SelectFields []string      // 查询字段
}

// GreetQueryResult 查询结果
type GreetQueryResult struct {
	Data       Greets
	PageResult *PaginationResult
}

// Greets 角色对象列表
type Greets []*Greet

// ToNames 获取角色名称列表
func (a Greets) ToNames() []string {
	names := make([]string, len(a))
	for i, item := range a {
		names[i] = item.Name
	}
	return names
}

// ToMap 转换为键值存储
func (a Greets) ToMap() map[uint64]*Greet {
	m := make(map[uint64]*Greet)
	for _, item := range a {
		m[item.ID] = item
	}
	return m
}

// ----------------------------------------GreetMenu--------------------------------------

// GreetMenu 角色菜单对象
type GreetMenu struct {
	ID       uint64 `json:"id,string"`                           // 唯一标识
	GreetID  uint64 `json:"Greet_id,string" binding:"required"`  // 角色ID
	MenuID   uint64 `json:"menu_id,string" binding:"required"`   // 菜单ID
	ActionID uint64 `json:"action_id,string" binding:"required"` // 动作ID
}

// GreetMenuQueryParam 查询条件
type GreetMenuQueryParam struct {
	PaginationParam
	GreetID  uint64   // 角色ID
	GreetIDs []uint64 // 角色ID列表
}

// GreetMenuQueryOptions 查询可选参数项
type GreetMenuQueryOptions struct {
	OrderFields  []*OrderField
	SelectFields []string
}

// GreetMenuQueryResult 查询结果
type GreetMenuQueryResult struct {
	Data       GreetMenus
	PageResult *PaginationResult
}

// GreetMenus 角色菜单列表
type GreetMenus []*GreetMenu

// ToMap 转换为map
func (a GreetMenus) ToMap() map[string]*GreetMenu {
	m := make(map[string]*GreetMenu)
	for _, item := range a {
		m[fmt.Sprintf("%d-%d", item.MenuID, item.ActionID)] = item
	}
	return m
}

// ToGreetIDMap 转换为角色ID映射
func (a GreetMenus) ToGreetIDMap() map[uint64]GreetMenus {
	m := make(map[uint64]GreetMenus)
	for _, item := range a {
		m[item.GreetID] = append(m[item.GreetID], item)
	}
	return m
}

// ToMenuIDs 转换为菜单ID列表
func (a GreetMenus) ToMenuIDs() []uint64 {
	var idList []uint64
	m := make(map[uint64]struct{})

	for _, item := range a {
		if _, ok := m[item.MenuID]; ok {
			continue
		}
		idList = append(idList, item.MenuID)
		m[item.MenuID] = struct{}{}
	}

	return idList
}

// ToActionIDs 转换为动作ID列表
func (a GreetMenus) ToActionIDs() []uint64 {
	idList := make([]uint64, len(a))
	m := make(map[uint64]struct{})
	for i, item := range a {
		if _, ok := m[item.ActionID]; ok {
			continue
		}
		idList[i] = item.ActionID
		m[item.ActionID] = struct{}{}
	}
	return idList
}
