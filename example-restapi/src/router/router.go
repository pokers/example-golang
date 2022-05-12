package router

import (
	"example-restapi/src/handler/user"
	"example-restapi/src/util"

	"github.com/gin-gonic/gin"
)

func RunRouter() {
	router := gin.Default()

	util.PrintInfo("regist handlers...")
	router.GET("api/:version/user", user.GetUserList)
	router.GET("api/:version/user/:id", user.GetUser)
	router.POST("api/:version/user", user.PostUser)

	util.PrintInfo("Run server...")
	router.Run(":8080")
}
