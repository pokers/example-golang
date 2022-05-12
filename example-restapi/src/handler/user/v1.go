package user

import (
	model "example-restapi/src/model"
	"example-restapi/src/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getUserListV1(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, model.UserList)
}

func getUserV1(ctx *gin.Context) {
	idParam, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "Not acceptable user id"})
		return
	}
	userId := uint(idParam)

	for _, user := range model.UserList {
		if user.Id == userId {
			ctx.IndentedJSON(http.StatusOK, user)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func postUserV1(ctx *gin.Context) {
	newUser := new(model.User)
	if err := ctx.BindJSON(newUser); err != nil {
		util.PrintError("%+v", err)
		ctx.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "Not acceptable user data"})
		return
	}

	util.PrintInfo("%+v", newUser)
	model.UserList = append(model.UserList, *newUser)
	ctx.IndentedJSON(http.StatusCreated, *newUser)
}
