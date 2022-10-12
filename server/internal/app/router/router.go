package router

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/mlsjla/gin-nuxt/server/internal/app/api"
	"github.com/mlsjla/gin-nuxt/server/internal/app/middleware"
	"github.com/mlsjla/gin-nuxt/server/internal/app/service"
	"github.com/mlsjla/gin-nuxt/server/pkg/auth"
)

var _ IRouter = (*Router)(nil)

var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))

type IRouter interface {
	Register(app *gin.Engine) error
	Prefixes() []string
}

type Router struct {
	Auth           auth.Auther
	CasbinEnforcer *casbin.SyncedEnforcer
	LoginAPI       *api.LoginAPI
	MenuAPI        *api.MenuAPI
	RoleAPI        *api.RoleAPI
	UserAPI        *api.UserAPI
	ThreadAPI      *api.ThreadAPI
	CategoryAPI    *api.CategoryAPI
	PostAPI        *api.PostAPI
	AttachmentAPI  *api.AttachmentAPI
	UploadAPI      *api.UploadAPI
	SettingAPI     *api.SettingAPI
	MobileCodeAPI  *api.MobileCodeAPI
	CasbinRuleAPI  *api.CasbinRuleAPI
	UserSrv        *service.UserSrv
	RoleMenuAPI    *api.RoleMenuAPI
	AppAPI         *api.AppAPI
	AppLogAPI      *api.AppLogAPI
} // end

func (a *Router) Register(app *gin.Engine) error {
	a.RegisterAPI(app)
	return nil
}

func (a *Router) Prefixes() []string {
	return []string{
		"/api/",
	}
}

// RegisterAPI register api group router
func (a *Router) RegisterAPI(app *gin.Engine) {
	g := app.Group("/api")

	// 游客允许的api，跳过接口验证
	allowGuest := []middleware.SkipperFunc{
		middleware.AllowPathPrefixSkipper("/api/v1/posts", "/api/v1/threads"),
	}
	g.Use(middleware.UserAuthMiddleware(a.Auth,
		allowGuest,
		middleware.AllowPathPrefixSkipper("/api/v1/pub/login", "/api/v1/pub/open"),
	))

	g.Use(middleware.CasbinMiddleware(a.CasbinEnforcer, a.UserSrv, allowGuest, middleware.AllowPathPrefixSkipper("/api/v1/pub")))

	g.Use(middleware.RateLimiterMiddleware())

	v1 := g.Group("/v1")
	{
		pub := v1.Group("/pub")
		{
			gLogin := pub.Group("login")
			{
				gLogin.GET("captchaid", a.LoginAPI.GetCaptcha)
				gLogin.GET("captcha", a.LoginAPI.ResCaptcha)
				gLogin.POST("", a.LoginAPI.Login)
				gLogin.POST("exit", a.LoginAPI.Logout)
			}

			gOpen := pub.Group("open")

			gOpen.GET("/categories.tree", a.CategoryAPI.QueryTree)
			gOpen.GET("/thread", a.ThreadAPI.Query)
			gOpen.GET("/thread/:id", a.ThreadAPI.Get)
			gOpen.GET("/post", a.PostAPI.Query)
			gOpen.GET("/post/:id", a.PostAPI.Get)
			gOpen.POST("wechat/:token", a.LoginAPI.WechatRegister)

			gCurrent := pub.Group("current")
			{
				gCurrent.POST("sendsms", a.LoginAPI.SendSms)
				gCurrent.PUT("password", a.LoginAPI.UpdatePassword)
				gCurrent.PUT("user", a.LoginAPI.UpdateUserInfo)
				gCurrent.GET("user", a.LoginAPI.GetUserInfo)
				gCurrent.GET("menutree", a.LoginAPI.QueryUserMenuTree)
			}
			pub.POST("/refresh-token", a.LoginAPI.RefreshToken)

		}

		gUpload := v1.Group("upload")
		{
			gUpload.POST("upload", a.UploadAPI.UploadFile)       // 上传文件
			gUpload.POST("getFileList", a.UploadAPI.GetFileList) // 获取上传文件列表
			gUpload.POST("deleteFile", a.UploadAPI.DeleteFile)   // 删除指定文件
			// gUpload.POST("breakpointContinue", a.UploadAPI.BreakpointContinue)             // 断点续传
			// gUpload.GET("findFile", a.UploadAPI.FindFile)                                  // 查询当前文件成功的切片
			// gUpload.POST("breakpointContinueFinish", a.UploadAPI.BreakpointContinueFinish) // 查询当前文件成功的切片
			// gUpload.POST("removeChunk", a.UploadAPI.RemoveChunk)                           // 查询当前文件成功的切片
		}

		gMenu := v1.Group("menus")
		{
			gMenu.GET("", a.MenuAPI.Query)
			gMenu.GET(":id", a.MenuAPI.Get)
			gMenu.POST("", a.MenuAPI.Create)
			gMenu.PUT(":id", a.MenuAPI.Update)
			gMenu.DELETE(":id", a.MenuAPI.Delete)
			gMenu.PATCH(":id/enable", a.MenuAPI.Enable)
			gMenu.PATCH(":id/disable", a.MenuAPI.Disable)
		}
		v1.GET("/menus.tree", a.MenuAPI.QueryTree)

		gRole := v1.Group("roles")
		{
			gRole.GET("", a.RoleAPI.Query)
			gRole.GET(":id", a.RoleAPI.Get)
			gRole.POST("", a.RoleAPI.Create)
			gRole.PUT(":id", a.RoleAPI.Update)
			gRole.DELETE(":id", a.RoleAPI.Delete)
			gRole.PATCH(":id/enable", a.RoleAPI.Enable)
			gRole.PATCH(":id/disable", a.RoleAPI.Disable)
		}
		v1.GET("/roles.select", a.RoleAPI.QuerySelect)

		gUser := v1.Group("users")
		{
			gUser.GET("", a.UserAPI.Query)
			gUser.GET(":id", a.UserAPI.Get)
			gUser.POST("", a.UserAPI.Create)
			gUser.PUT(":id", a.UserAPI.Update)
			gUser.DELETE(":id", a.UserAPI.Delete)
			gUser.PATCH(":id/enable", a.UserAPI.Enable)
			gUser.PATCH(":id/disable", a.UserAPI.Disable)
		}

		gThread := v1.Group("threads")
		{
			gThread.GET("", a.ThreadAPI.Query)
			gThread.GET(":id", a.ThreadAPI.Get)
			gThread.POST("", a.ThreadAPI.Create)
			gThread.PUT(":id", a.ThreadAPI.Update)
			gThread.DELETE(":id", a.ThreadAPI.Delete)

		}

		gCategory := v1.Group("categories")
		{
			gCategory.GET("", a.CategoryAPI.Query)
			gCategory.GET(":id", a.CategoryAPI.Get)
			gCategory.POST("", a.CategoryAPI.Create)
			gCategory.PUT(":id", a.CategoryAPI.Update)
			gCategory.DELETE(":id", a.CategoryAPI.Delete)

		}
		v1.GET("/categories.tree", a.CategoryAPI.QueryTree)

		gPost := v1.Group("posts")
		{
			gPost.GET("", a.PostAPI.Query)
			gPost.GET(":id", a.PostAPI.Get)
			gPost.POST("", a.PostAPI.Create)
			gPost.PUT(":id", a.PostAPI.Update)
			gPost.DELETE(":id", a.PostAPI.Delete)

		}
		v1.GET("/posts.tree", a.PostAPI.QueryTree)

		gAttachment := v1.Group("attachments")
		{
			gAttachment.GET("", a.AttachmentAPI.Query)
			gAttachment.GET(":id", a.AttachmentAPI.Get)
			gAttachment.POST("", a.AttachmentAPI.Create)
			gAttachment.PUT(":id", a.AttachmentAPI.Update)
			gAttachment.DELETE(":id", a.AttachmentAPI.Delete)

		}

		gSetting := v1.Group("settings")
		{
			gSetting.GET("", a.SettingAPI.Query)
			gSetting.GET(":id", a.SettingAPI.Get)
			gSetting.POST("", a.SettingAPI.Create)
			gSetting.PUT(":id", a.SettingAPI.Update)
			gSetting.DELETE(":id", a.SettingAPI.Delete)

		}

		gMobileCode := v1.Group("mobile-codes")
		{
			gMobileCode.GET("", a.MobileCodeAPI.Query)
			gMobileCode.GET(":id", a.MobileCodeAPI.Get)
			gMobileCode.POST("", a.MobileCodeAPI.Create)
			gMobileCode.PUT(":id", a.MobileCodeAPI.Update)
			gMobileCode.DELETE(":id", a.MobileCodeAPI.Delete)

		}

		gCasbinRule := v1.Group("casbin-rules")
		{
			gCasbinRule.GET("", a.CasbinRuleAPI.Query)
			gCasbinRule.GET(":id", a.CasbinRuleAPI.Get)
			gCasbinRule.POST("", a.CasbinRuleAPI.Create)
			gCasbinRule.PUT(":id", a.CasbinRuleAPI.Update)
			gCasbinRule.DELETE(":id", a.CasbinRuleAPI.Delete)
			gCasbinRule.GET("api", a.CasbinRuleAPI.QueryUseApi)

		}

		gRoleMenu := v1.Group("role-menus")
		{
			gRoleMenu.GET("", a.RoleMenuAPI.Query)
			gRoleMenu.GET(":id", a.RoleMenuAPI.Get)
			gRoleMenu.POST("", a.RoleMenuAPI.Create)
			gRoleMenu.PUT(":id", a.RoleMenuAPI.Update)
			gRoleMenu.DELETE(":id", a.RoleMenuAPI.Delete)

		}

		gApp := v1.Group("apps")
		{
			gApp.GET("", a.AppAPI.Query)
			gApp.GET(":id", a.AppAPI.Get)
			gApp.POST("", a.AppAPI.Create)
			gApp.PUT(":id", a.AppAPI.Update)
			gApp.DELETE(":id", a.AppAPI.Delete)

		}

		gAppLog := v1.Group("app-logs")
		{
			gAppLog.GET("", a.AppLogAPI.Query)
			gAppLog.GET(":id", a.AppLogAPI.Get)
			gAppLog.POST("", a.AppLogAPI.Create)
			gAppLog.PUT(":id", a.AppLogAPI.Update)
			gAppLog.DELETE(":id", a.AppLogAPI.Delete)

		}
		v1.POST("/app-logs.preview/:id", a.AppLogAPI.Preview)
		v1.POST("/app-logs.upload/:id", a.AppLogAPI.Upload)

	} // v1 end
}
