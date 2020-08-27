package main

import (
	"blogger/controller"
	"blogger/dao/db"
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	dns := "root:root@tcp(localhost:3306)/blogger?parseTime=true"
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}
	ginpprof.Wrapper(router)
	router.Static("/static/", "./static")
	router.LoadHTMLGlob("views/**/*")
	router.GET("/", controller.IndexHandle)
	router.GET("/category/", controller.CategoryList)
	router.GET("/article/new/", controller.NewArticle)
	router.POST("/article/submit/", controller.ArticleSubmit)
	router.GET("/article/detail/", controller.ArticleDetail)
	router.POST("/upload/file/", controller.UploadFile)
	router.GET("/leave/new/", controller.LeaveNew)
	router.GET("/about/me/", controller.AboutMe)
	router.POST("/comment/submit/", controller.CommentSubmit)
	router.POST("/leave/submit/", controller.LeaveSubmit)
    //后台路由器
	router.GET("/back/login", controller.BackLoginHandle)
	router.GET("/back/index", controller.BackIndexHandle)
	router.POST("/back/dologin", controller.BackDoLoginHandle)
	router.GET("/GetCaptcha", controller.GetCaptcha)
	router.GET("/verifyCaptcha", controller.VerifyCaptcha)
	router.GET("/show/:source", controller.GetCaptchaPng)
	_ = router.Run(":8010")
}
