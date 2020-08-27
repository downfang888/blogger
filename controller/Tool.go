package controller

import (
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

func GetCaptchass(c *gin.Context) {
	captcha.Server(captcha.StdWidth, captcha.StdHeight)
}