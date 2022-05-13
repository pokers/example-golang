package router

import (
	"example-restapi+db/src/handler/user"
	"example-restapi+db/src/util"

	"github.com/gin-gonic/gin"
)

func RunRouter() {
	router := gin.Default()

	util.PrintInfo("regist handlers...")
	userHandlerV1 := user.UserHandlerV1{}
	routerV1 := router.Group("/api/v1/")
	{
		routerV1.GET("/user", userHandlerV1.GetUserList)
		routerV1.GET("/user/:id", userHandlerV1.GetUser)
		routerV1.POST("/user", userHandlerV1.PostUser)
	}

	userHandlerV2 := user.UserHandlerV2{}
	routerV2 := router.Group("/api/v2")
	{
		routerV2.GET("/user", userHandlerV2.GetUserList)
		routerV2.GET("/user/:id", userHandlerV2.GetUser)
		routerV2.POST("/user", userHandlerV2.PostUser)
	}

	util.PrintInfo("Run server...")
	router.Run(":8080")
}
