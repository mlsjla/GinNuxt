package setting

import (
	"context"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
)

// RouteInfo represents a request route's specification which contains method and path and its handler.
type RouteInfo struct {
	Method  string `json:"method"`
	Path    string `json:"path"`
	Handler string `json:"-"`
}

// RoutesInfo defines a RouteInfo array.
type RoutesInfo []RouteInfo

func (l RoutesInfo) Len() int {
	return len(l)
}
func (l RoutesInfo) Less(i, j int) bool {
	return l[i].Path < l[j].Path
}
func (l RoutesInfo) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

var (
	S       *schema.SettingQueryResult
	Routers RoutesInfo
	// once sync.Once
)

func InitSetting(ctx context.Context, settingResult *schema.SettingQueryResult) error {
	S = settingResult
	return nil
}
func StoreRouters(ctx context.Context, routes gin.RoutesInfo) error {
	var data RoutesInfo
	for _, r := range routes {
		data = append(data, RouteInfo{
			Method:  r.Method,
			Path:    r.Path,
			Handler: r.Handler,
		})
	}
	sort.Sort(data)
	Routers = data
	return nil
}

func Setting(key string) string {
	value := ""
	for i := 0; i < len(S.Data); i++ {
		if key == S.Data[i].Key {
			value = S.Data[i].Value
			break
		}
	}
	return value
}
