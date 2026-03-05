package controller

import (
	"github.com/gin-gonic/gin"
)

type CommonController struct {
}

func (s *CommonController) TestFun(c *gin.Context) {
	msg := "this is TestFun"
	println(msg)
	return
}
