package user

import (
	model "example-restapi+db/src/model"
	"example-restapi+db/src/service/sql"
	"example-restapi+db/src/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandlerV1 struct{}

func (User *UserHandlerV1) GetUserList(ctx *gin.Context) {
	var user []model.User
	inst := sql.GetConnection()
	util.PrintInfo("DB Inst : %+v", inst)

	if err := inst.Sql.Debug().Table("user").Limit(10).Find(&user).Error; err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server Error"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, user)
}

func (User *UserHandlerV1) GetUser(ctx *gin.Context) {
	idParam, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "Not acceptable user id"})
		return
	}
	userId := uint(idParam)

	var user model.User
	inst := sql.GetConnection()
	util.PrintInfo("DB Inst : %+v", inst)

	if err := inst.Sql.Debug().Table("user").Find(&user, userId).Error; err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, user)
}

func (User *UserHandlerV1) PostUser(ctx *gin.Context) {
	newUser := new(model.User)
	if err := ctx.BindJSON(newUser); err != nil {
		util.PrintError("%+v", err)
		ctx.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "Not acceptable user data"})
		return
	}

	inst := sql.GetConnection()
	util.PrintInfo("DB Inst : %+v", inst)
	util.PrintInfo("%+v", newUser)

	if err := inst.Sql.Debug().Table("user").Create(&newUser).Error; err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server Error"})
		return
	}
	ctx.IndentedJSON(http.StatusCreated, *newUser)
}
