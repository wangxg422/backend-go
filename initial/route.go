package initial

import (
	"backend/global"
	"backend/route/public"
	"backend/route/system"

	"github.com/gin-gonic/gin"
)

func InitRoute() {
	// debug or release
	gin.SetMode(global.AppConfig.App.Mode)
	r := gin.Default()

	// 开放接口，健康检查、注册、忘记密码等
	publicRouterGroup := &public.PublicRouterGroup{}
	publicGroup := r.Group("public")
	{
		publicRouterGroup.AddPublicRouterGroup(publicGroup)
	}

	sysRouterGroup := &system.SysUserRouter{}
	sysGroup := r.Group("sys")
	{
		sysRouterGroup.AddSysUserRouter(sysGroup)
	}

	_ = r.Run(":" + global.AppConfig.App.Port)
}
