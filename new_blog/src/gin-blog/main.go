package main

import (
	"fmt"
	"net/http"


	"new_blog/src/gin-blog/routers"
	"new_blog/src/gin-blog/pkg/setting"
)
func main() {
	//router :=gin.Default()
	//router.GET("/test", func(context *gin.Context) {
	//	context.JSON(200,gin.H{
	//		"messgae":"test",
	//	})
	//})
	router :=routers.InitRouter()

	s :=&http.Server{
		Addr: fmt.Sprintf(":%d",setting.HTTPPort),
		Handler:   router,
		ReadTimeout: setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
		MaxHeaderBytes: 1<<20,

	}
	s.ListenAndServe()
}
