package handler

import (
	util "example-restapi+db/src/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlePanic(ctx *gin.Context) {
	if exception := recover(); exception != nil {
		util.PrintError("Panic!!")
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"exception": "Internal server error..."})
	}
}
