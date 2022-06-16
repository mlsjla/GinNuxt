package schema

import "time"

// 文章分类
type Category struct {
	ID          uint64          `json:"id,string"`
	Name        string          `json:"name" binding:"required"` // 分类名称
	Description string          `json:"description"`             // 分类描述
	Icon        string          `json:"icon"`                    // 图标
	Sort        int             `json:"sort"`                    // 排序
	Property    int             `json:"property"`                // 属性
	ThreadCount uint64          `json:"thread_count,string"`     // 主题数
	Moderators  string          `json:"moderators"`              // 排序
	Ip          string          `json:"ip"`                      // 排序
	Parentid    uint64          `json:"Parentid,string"`         // 父ID
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	Actions     CategoryActions `json:"actions"`
}

// Query parameters for db
type CategoryQueryParam struct {
	PaginationParam
}

// Query options for db (order or select fields)
type CategoryQueryOptions struct {
	OrderFields  []*OrderField
	SelectFields []string
}

// Query result from db
type CategoryQueryResult struct {
	Data       Categories
	PageResult *PaginationResult
}

// CategoryActionResource 分类动作关联资源对象
type CategoryActionResource struct {
	ID       uint64 `yaml:"-" json:"id,string"`                      // 唯一标识
	ActionID uint64 `yaml:"-" json:"action_id,string"`               // 分类动作ID
	Method   string `yaml:"method" binding:"required" json:"method"` // 资源请求方式(支持正则)
	Path     string `yaml:"path" binding:"required" json:"path"`     // 资源请求路径（支持/:id匹配）
}

// CategoryActionResources 分类动作关联资源管理列表
type CategoryActionResources []*CategoryActionResource

// CategoryAction 分类动作对象
type CategoryAction struct {
	ID         uint64                  `yaml:"-" json:"id,string"`                             // 唯一标识
	CategoryID uint64                  `yaml:"-" binding:"required" json:"category_id,string"` // 分类ID
	Code       string                  `yaml:"code" binding:"required" json:"code"`            // 动作编号
	Name       string                  `yaml:"name" binding:"required" json:"name"`            // 动作名称
	Resources  CategoryActionResources `yaml:"resources,omitempty" json:"resources"`           // 资源列表
}

// CategoryActionResourceQueryParam 查询条件
type CategoryActionResourceQueryParam struct {
	PaginationParam
	CategoryID  uint64   // 分类ID
	CategoryIDs []uint64 // 分类ID列表
}

// 文章分类 Object List
type Categories []*Category

// CategoryActions 分类动作管理列表
type CategoryActions []*CategoryAction

type CategoryTree struct {
	ID          uint64          `json:"id,string"`
	Name        string          `json:"name" binding:"required"` // 分类名称
	Description string          `json:"description"`             // 分类描述
	Icon        string          `json:"icon"`                    // 图标
	Sort        int             `json:"sort"`                    // 排序
	Property    int             `json:"property"`                // 属性
	ThreadCount uint64          `json:"thread_count"`            // 主题数
	Moderators  string          `json:"moderators"`              // 排序
	Ip          string          `json:"ip"`                      // 排序
	Parentid    uint64          `json:"Parentid"`                // 父ID
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`                                   // 排序值
	Actions     CategoryActions `yaml:"actions,omitempty" json:"actions"`             // 动作列表
	Children    *CategoryTrees  `yaml:"children,omitempty" json:"children,omitempty"` // 子级树
}

// CategoryTrees 分类树列表
type CategoryTrees []*CategoryTree

// ToTree 转换为树形结构
func (a CategoryTrees) ToTree() CategoryTrees {
	mi := make(map[uint64]*CategoryTree)
	for _, item := range a {
		mi[item.ID] = item
	}

	var list CategoryTrees
	for _, item := range a {
		if item.Parentid == 0 {
			list = append(list, item)
			continue
		}
		if pitem, ok := mi[item.Parentid]; ok {
			if pitem.Children == nil {
				children := CategoryTrees{item}
				pitem.Children = &children
				continue
			}
			*pitem.Children = append(*pitem.Children, item)
		}
	}
	return list
}

// ToTree 转换为分类树
// ToTree 转换为菜单树
func (a Categories) ToTree() CategoryTrees {
	list := make(CategoryTrees, len(a))
	for i, item := range a {
		list[i] = &CategoryTree{
			ID:          item.ID,
			Name:        item.Name,
			Icon:        item.Icon,
			ThreadCount: item.ThreadCount,
			Parentid:    item.Parentid,
			Sort:        item.Sort,
			Actions:     item.Actions,
		}
	}
	return list.ToTree()
}
