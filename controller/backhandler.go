package controller

import (
	"blogger/service"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)
//
func BackIndexHandle(c *gin.Context) {
	c.HTML(http.StatusOK, "views/back/index.html", gin.H{})
}
func BackLoginHandle(c *gin.Context) {
	c.HTML(http.StatusOK, "views/back/login.html", gin.H{})
}
//处理登录验证
func BackDoLoginHandle(c *gin.Context) {
	username := c.PostForm("loginuser")
	password := c.PostForm("loginpwd")
	err := service.GetLogin(username,password)
	//返回数据给浏览器
	response := gin.H{}
	if err == nil {
		session := sessions.Default(c)
		session.Set("adminlogin-1", username)
	err :=	session.Save()
		if err != nil {
			fmt.Printf("Fail to save session--%v",err)
		} else {
			fmt.Println("Succeed to save session")
		}
		//无误
		response = gin.H{"code": 200, "message": "ok----aaaa"}
	} else {
		response = gin.H{"code": 101, "message": "error-----2222"}
	}
	c.JSON(http.StatusOK, response)
	//return
}