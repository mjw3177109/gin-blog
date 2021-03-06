package routers

import(
	"github.com/gin-gonic/gin"
	"new_blog/src/gin-blog/pkg/setting"
	"new_blog/src/gin-blog/routers/api/v1"
)

func InitRouter()*gin.Engine{
	r :=gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 :=r.Group("/api/v1s")
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message":"test",
		})
	})
	return r
}