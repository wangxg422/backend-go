package system

import (
	"github.com/gin-gonic/gin"
)

type SysPostRouter struct {
}

func (m *SysPostRouter) AddSysPostRouter(route *gin.RouterGroup) {
	router := route.Group("post")
	{
		router.POST("", postApi.CreatePost)
		router.GET("list", postApi.ListPost)
		router.GET(":postId", postApi.GetPostById)
		router.POST("update", postApi.UpdatePost)
		router.DELETE(":postId", postApi.DeletePost)
	}
}
