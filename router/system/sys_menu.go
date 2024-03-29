package system

import (
	"github.com/gin-gonic/gin"
)

type SysMenuRouter struct {
}

func (m *SysMenuRouter) AddSysMenuRouter(route *gin.RouterGroup) {
	router := route.Group("menu")
	{
		router.POST("", menuApi.CreateMenu)
		router.GET("list", menuApi.ListMenu)
		router.GET(":deptId", menuApi.GetMenuById)
		router.POST("update", menuApi.UpdateMenu)
		router.DELETE(":deptId", menuApi.DeleteMenu)
	}
}
