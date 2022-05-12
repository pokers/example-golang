package user

import (
	"net/http"

	"example-restapi/src/handler"

	"github.com/gin-gonic/gin"
)

func GetUserList(ctx *gin.Context) {
	defer handler.HandlePanic(ctx)
	version := ctx.Param("version")

	switch version {
	case "v1":
		getUserListV1(ctx)
	case "v2":
		getUserListV2(ctx)
	default:
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	}
}

func GetUser(ctx *gin.Context) {
	defer handler.HandlePanic(ctx)
	version := ctx.Param("version")

	switch version {
	case "v1":
		getUserV1(ctx)
	case "v2":
		getUserV2(ctx)
	default:
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	}
}

func PostUser(ctx *gin.Context) {
	defer handler.HandlePanic(ctx)
	version := ctx.Param("version")

	switch version {
	case "v1":
		postUserV1(ctx)
	case "v2":
		postUserV2(ctx)
	default:
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	}
}
