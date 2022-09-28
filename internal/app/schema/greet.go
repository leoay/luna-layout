package schema

// Greet 角色对象
type Greet struct {
	ID   uint64 `json:"id,string"`               // 唯一标识
	Name string `json:"name" binding:"required"` // 角色名称
}

// GreetQueryParam 查询条件
type GreetQueryParam struct {
	PaginationParam
	IDs        []uint64 `form:"-"`          // 唯一标识列表
	Name       string   `form:"-"`          // 角色名称
	QueryValue string   `form:"queryValue"` // 模糊查询
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
