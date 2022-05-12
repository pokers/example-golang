package user

import (
	"example-restapi/src/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getUserListV2(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, model.UserProfileList)
}

func getUserV2(ctx *gin.Context) {
	idParam, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "Not acceptable user id"})
		return
	}
	userId := uint(idParam)
	for _, user := range model.UserProfileList {
		if user.UserId == userId {
			ctx.IndentedJSON(http.StatusOK, user)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func postUserV2(ctx *gin.Context) {
	newUser := new(model.UserProfile)
	if err := ctx.BindJSON(newUser); err != nil {
		ctx.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "Not acceptable user data"})
		return
	}

	model.UserProfileList = append(model.UserProfileList, *newUser)
	ctx.IndentedJSON(http.StatusCreated, *newUser)
}
